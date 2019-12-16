package parser

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"regexp"
)

// cityListRe城市列表匹配正则
var (
	cityRe   = regexp.MustCompile(`<a\s+href="(http://album.zhenai.com/u/\d+)"[^>]+>([^<]+)</a>`)
	genderRe = regexp.MustCompile(`<td\s+width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
)

// ParseCity解析HTTP响应内容城市
func ParseCity(contents []byte) (parseResult engine.ParseResult) {
	cities := UsersMergeOfCity(contents)
	for _, city := range cities {
		name := city.Name
		gender := city.Gender
		parseResult.Items = append(parseResult.Items, name)
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: string(city.URL),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name, gender)
			},
		})
	}
	return
}

// usersMergeOfCity合并用户链接及性别
func UsersMergeOfCity(contents []byte) model.Cities {
	cities := cityRe.FindAllSubmatch(contents, -1)
	genders := genderRe.FindAllSubmatch(contents, -1)
	r := model.Cities{}
	for index, city := range cities {
		for i, gender := range genders {
			if i == index {
				c := model.City{}
				c.Name = string(city[2])
				c.URL = string(city[1])
				c.Gender = string(gender[1])
				r = append(r, c)
			}
		}
	}
	return r
}
