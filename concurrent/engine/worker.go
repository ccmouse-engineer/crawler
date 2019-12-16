package engine

import (
	"crawler/concurrent/fetcher"

	"github.com/sirupsen/logrus"
)

// worker发送HTTP请求并解析响应数据
func worker(r Request) (ParseResult, error) {
	// 发送HTTP请求并获取响应数据
	resp, err := fetcher.Fetch(r.Url)
	if err != nil {
		logrus.Errorf("Fetch: error fetching url: %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	// 调用解析器进行解析响应数据
	return r.ParserFunc(resp), nil
}
