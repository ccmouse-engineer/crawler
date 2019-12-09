package engine

import (
	"crawler/fetcher"

	"github.com/sirupsen/logrus"
)

// SimpleEngine简单引擎
type SimpleEngine struct{}

// Run爬虫入口
func (s SimpleEngine) Run(seeds ...Request) {
	// 请求队列
	var requests []Request

	// 请求放入到请求队列
	for _, req := range seeds {
		requests = append(requests, req)
	}

	// 只要请求队列不为空就继续
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		// worker发送HTTP请求并解析响应数据
		parseResult, err := worker(req)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			logrus.Infof("Got item: %v", item)
		}
	}
}

// worker发送HTTP请求并解析响应数据
func worker(r Request) (ParseResult, error) {
	// 发送HTTP请求并获取响应数据
	resp, err := fetcher.Fetch(r.Url)
	// log.Printf("Fetching url: %s", r.Url)
	if err != nil {
		logrus.Errorf("Fetch: error fetching url: %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// 调用解析器进行解析响应数据
	return r.ParserFunc(resp), nil
}
