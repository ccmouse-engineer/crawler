package parser

import (
	"crawler/engine"
	"regexp"
)

// cityListRe城市列表匹配正则
const cityListRe = `<a\s+href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`

// ParseCityList解析HTTP响应内容城市列表
func ParseCityList(contents []byte) (parseResult engine.ParseResult) {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, match := range matchs {
		parseResult.Items = append(parseResult.Items, string(match[2]))
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:        string(match[1]),
			ParserFunc: ParseCity,
		})
	}
	return
}
