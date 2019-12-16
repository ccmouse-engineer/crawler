package parser

import (
	"crawler/concurrent/engine"
	"regexp"
	"strings"
)

// cRe车型列表匹配正则
var (
	modelRe = regexp.MustCompile(`<a\s+href="(/\w+/)"\s+target="_blank"\s+title="[^"]+">([^>]+[^<]+)</a>`)
)

// ParseCar解析HTTP响应内容车系
func ParseModel(contents []byte, domain string, model string) (parseResult engine.ParseResult) {
	cars := modelRe.FindAllSubmatch(contents, -1)
	for _, car := range cars {
		url := domain + strings.TrimSpace(string(car[1]))
		name := string(car[2])
		parseResult.Items = append(parseResult.Items, name)
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseDetail(c, name, model)
			},
		})
	}
	return
}
