package engine

import (
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
