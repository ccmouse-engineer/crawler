package parser

import (
	"crawler/engine"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// 正则表达式
var (
	// 匹配城市页面下的二手房列表项
	houseRe = regexp.MustCompile(`<a[\s\w="-]+href="([\w:/.]*html)"[\s\w="-]+data-housecode="\w+"[\s\w="-]+>([^><]+)</a>`)
	// 匹配城市页面下的二手房更多页
	nextRe = regexp.MustCompile(`<div\s+class="page-box\s+house-lst-page-box"\s+comp-module='page'\s+page-url="/ershoufang/pg\{page\}"page-data='\{"totalPage":(\d+),"curPage":(\d+)\}'></div>`)
	// 匹配城市页面二手房详情URL
	idURLRe = regexp.MustCompile(`[\w:/.]+[/]{1}(\d+).html`)
)

// ParseCity解析HTTP响应内容城市
func ParseCity(contents []byte, cityName, prefixURL string, isFirstPage bool) (parseResult engine.ParseResult) {
	matchsHouse := houseRe.FindAllSubmatch(contents, -1)
	houses := LinkDeduplication(matchsHouse)
	for _, house := range houses {
		for url, title := range house {
			match := idURLRe.FindSubmatch([]byte(url))
			id := string(match[1])
			parseResult.Requests = append(parseResult.Requests, engine.Request{
				Url: url,
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseDetail(c, id, url, title, cityName)
				},
			})
		}
	}

	if isFirstPage {
		matchsNext := nextRe.FindSubmatch(contents)
		pages := generateMultiPages(matchsNext, prefixURL)
		for _, page := range pages {
			for url, title := range page {
				parseResult.Requests = append(parseResult.Requests, engine.Request{
					Url:        url,
					ParserFunc: CityParse(cityName+"-"+title, prefixURL, false),
				})
			}
		}

		curCityURLs := findMoreCurrentCityURLs(contents, prefixURL)
		for _, curCityURL := range curCityURLs {
			for url, title := range curCityURL {
				log.Printf("city #%s, URL: #%s\n", title, url)
				parseResult.Requests = append(parseResult.Requests, engine.Request{
					Url:        url,
					ParserFunc: CityParse(cityName+"-"+title, prefixURL, false),
				})
			}
		}
	}

	return
}

// CityParse城市解析
func CityParse(name, url string, isFirstPage bool) engine.ParserFunc {
	return func(contents []byte) engine.ParseResult {
		return ParseCity(contents, name, url, isFirstPage)
	}
}

// findMoreCurrentCityURLs获取当前城市下更多的区域二手房链接
func findMoreCurrentCityURLs(contents []byte, prefixURL string) []map[string]string {
	// 匹配城市页面下的二手房更多区域
	var moreRe = regexp.MustCompile(fmt.Sprintf(`<a\s+target="_blank"\s+href="(%s\w+/)">([^<]+)</a>`, prefixURL))
	var matches = moreRe.FindAllSubmatch(contents, -1)
	var result = make([]map[string]string, 0)
	for _, match := range matches {
		url := string(match[1])
		name := string(match[2])
		result = append(result, map[string]string{url: name})
	}
	return result
}

// generateMultiPages生成每个二手房城市下的更新页面
func generateMultiPages(matchs [][]byte, url string) []map[string]string {
	if len(matchs) > 0 && len(matchs) == 3 {
		totalPageNum := string(matchs[1])
		curPageNum := string(matchs[2])
		total, _ := strconv.Atoi(totalPageNum)
		cur, _ := strconv.Atoi(curPageNum)
		result := make([]map[string]string, 0)
		for i := cur + 1; i <= total; i++ {
			index := strconv.Itoa(i)
			result = append(result, map[string]string{url + "pg" + index: "第" + index + "页"})
		}
		return result
	}
	return nil
}
