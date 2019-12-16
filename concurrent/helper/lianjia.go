package helper

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

// GetFullDomainURL返回分类全URL地址
func GetFullDomainURL(domain, category string) (url string) {
	return domain + category + "/"
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

// MustPrice价格处理
func MustPrice(str string) (price float64) {
	if len(str) == 0 {
		return
	}
	price, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalf("must price error: %s\n", err)
		return
	}
	return
}

// GenerateMultiPages生成每个二手房城市下的更新页面
func GenerateMultiPages(matchs [][]byte, url string) []string {
	if len(matchs) > 0 && len(matchs) == 3 {
		totalPageNum := string(matchs[1])
		curPageNum := string(matchs[2])
		total, _ := strconv.Atoi(totalPageNum)
		cur, _ := strconv.Atoi(curPageNum)
		result := make([]string, 0)
		for i := cur + 1; i <= total; i++ {
			index := strconv.Itoa(i)
			result = append(result, url+"pg"+index)
		}
		return result
	}
	return nil
}
