package engine

import (
	"crawler/fetcher"

	"github.com/sirupsen/logrus"
)

// Run爬虫入口
func Run(seeds ...Request) {
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

		// 发送HTTP请求并获取响应数据
		resp, err := fetcher.Fetch(req.Url)
		logrus.Infof("Fetching url: %s", req.Url)
		if err != nil {
			logrus.Errorf("Fetch: error fetching url: %s: %v", req.Url, err)
			continue
		}

		// 调用解析器进行解析响应数据
		parseResult := req.ParserFunc(resp)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			logrus.Infof("Got item: %v", item)
		}
	}
}
