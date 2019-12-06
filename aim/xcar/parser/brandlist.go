package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

// brandListRe城市列表匹配正则
const brandListRe = `<li><a href="(/car/[0-9-]+/)"\s+[^>]+><span class="sign"><img src="//img1.xcarimg.com/PicLib/logo/[^"]+"><\/span>([^>]+[^<]+)</a></li>`

// ParseBrandList解析HTTP响应内容城市列表
func ParseBrandList(contents []byte, domain string) (parseResult engine.ParseResult) {
	re := regexp.MustCompile(brandListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, match := range matchs {
		url := domain + string(match[1])
		brand := strings.TrimSpace(string(match[2]))
		parseResult.Items = append(parseResult.Items, brand)
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseCar(c, domain, brand)
			},
		})
	}
	return
}
