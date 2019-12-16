package parser

import (
	"crawler/concurrent/engine"
	"regexp"
	"strings"
)

// cRe车系列表匹配正则
var (
	cRe = regexp.MustCompile(`<a\s+href="(/car/select/\w+/)"\s+target="_blank"\s+title="[^"]+">([^>]+[^<]+)</a>`)
)

// ParseCar解析HTTP响应内容车系
func ParseCar(contents []byte, domain string, brand string) (parseResult engine.ParseResult) {
	cars := cRe.FindAllSubmatch(contents, -1)
	for _, car := range cars {
		name := brand + "|" + string(car[2])
		url := domain + strings.TrimSpace(string(car[1]))
		parseResult.Items = append(parseResult.Items, name)
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseModel(c, domain, name)
			},
		})
	}
	return
}
