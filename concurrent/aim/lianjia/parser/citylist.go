package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/helper"
	"regexp"
	"sort"
	"strings"
)

// cityListRe城市列表匹配正则
var cityListRe = regexp.MustCompile(`<a\s+href="(https://\w+.lianjia.com/)">([^<]+)</a>`)

// ParseCityList解析HTTP响应内容城市列表
func ParseCityList(contents []byte, category string, limit int) (parseResult engine.ParseResult) {
	matchs := cityListRe.FindAllSubmatch(contents, -1)
	cities := LinkDeduplication(matchs)
	start := 1
	for _, city := range cities {
		if limit != -1 && start > limit {
			break
		}
		start++
		for url, name := range city {
			url = helper.GetFullDomainURL(url, category)
			parseResult.Requests = append(parseResult.Requests, engine.Request{
				Url:        url,
				ParserFunc: CityParse(name, url, true),
			})
		}
	}
	return
}

// LinkDeduplication连接列表去重
func LinkDeduplication(matchs [][][]byte) []map[string]string {
	var links = make(map[string]string, 0)
	for _, match := range matchs {
		if match == nil {
			continue
		}
		url := strings.TrimSpace(string(match[1]))
		name := strings.TrimSpace(string(match[2]))
		links[url] = name
	}

	var urls []string
	for url := range links {
		urls = append(urls, url)
	}

	sort.Slice(urls, func(i, j int) bool {
		return urls[i] < urls[j]
	})

	var res []map[string]string
	for _, url := range urls {
		city, ok := links[url]
		if ok {
			res = append(res, map[string]string{url: city})
		}
	}

	return res
}
