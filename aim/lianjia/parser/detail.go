package parser

import (
	"bytes"
	"crawler/engine"
	"crawler/helper"
	"crawler/model"
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	doorModelFigureWidth  = 960 // 户型图宽度
	doorModelFigureHeight = 640 // 户型图高度
)

var (
	// 总价
	priceRe = regexp.MustCompile(`<span\s+class="total">([^<>]+)</span>`)
	// 每平单价
	unitPriceRe = regexp.MustCompile(`<span\s+class="unitPriceValue">([^<>]+)<i>[^<>]+</i></span>`)
	// 小区名称
	communityNameRe = regexp.MustCompile(`<div\s+class="communityName">[<>/\w\s="]*<span\s+class="label">小区名称</span>[<>/\w\s="]*<a[/\w\s="]*>([^<>]+)</a>`)
	// 所在区域
	areaNameRe = regexp.MustCompile(`<div\s+class="areaName">[<>/\w\s="]*<span\s+class="label">所在区域</span>[<>/\w\s="]*([^<>]+)</a>[\w&;\s]*<a[/\w\s="]*>([^<>]+)</a>`)
	// 房源照片
	imageRe = regexp.MustCompile(`<div[\s\w-="]+>\s+<img src="([\w/:.-]+)"\s+alt="[^"]+">\s+<span\sclass="name">([^<]+)</span>\s+</div>`)
)

// ParseDetail解析HTTP响应内容二手房详情页
func ParseDetail(contents []byte, title string) engine.ParseResult {
	ershoufang, err := parseHouseBasicInfo(&contents, title)
	if err != nil {
		panic(err)
	}
	ershoufang.Title = title
	ershoufang.CommunityName = parseHouseFields(&contents, communityNameRe)
	ershoufang.AreaName = parseHouseFields(&contents, areaNameRe)
	ershoufang.Price = helper.MustPrice(parseHouseFields(&contents, priceRe))
	ershoufang.UnitPrice = helper.MustPrice(parseHouseFields(&contents, unitPriceRe))
	ershoufang.Images = parseHouseImages(&contents, imageRe)
	ershoufang.Characteristics = parseHouseCharacteristics(&contents)
	doorModelBetweenPoints, err := parseDoorModelBetweenPoints(&contents)
	if err != nil {
		panic(err)
	}
	ershoufang.DoorModelBetweenPoints = *doorModelBetweenPoints
	return engine.ParseResult{Items: []interface{}{ershoufang}}
}

// parseDoorModelBetweenPoints获取二手房详情页面户型分间数据
func parseDoorModelBetweenPoints(contents *[]byte) (*model.ErshoufangDoorModelBetweenPoints, error) {
	reader := bytes.NewReader(*contents)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	var doorModelBetweenPoints = new(model.ErshoufangDoorModelBetweenPoints)
	var roomData = make([]map[string][]string, 0)
	var doorModelFigure = ""
	doc.Find("#layout div .content").Each(func(index int, selection *goquery.Selection) {
		// 户型图
		img, exists := selection.Find(".imgdiv").Attr("data-img")
		if exists {
			filename := fmt.Sprintf(".%dx%d.jpg", doorModelFigureWidth, doorModelFigureHeight)
			doorModelFigure = img + filename
		}

		// 户型房间数据
		selection.Find(".des .info .list #infoList").Each(func(ii int, is *goquery.Selection) {
			is.Children().Each(func(ri int, rs *goquery.Selection) {
				var data = make([]string, 0)
				var title = ""
				rs.Find("div").Each(func(di int, ds *goquery.Selection) {
					var content = strings.TrimSpace(ds.Text())
					if di == 0 {
						title = content
					} else {
						data = append(data, content)
					}

				})
				roomData = append(roomData, map[string][]string{title: data})
			})
		})
	})
	doorModelBetweenPoints.RoomData = roomData
	doorModelBetweenPoints.DoorModelFigure = doorModelFigure
	return doorModelBetweenPoints, nil
}

// parseHouseImages获取二手房详情页面房源照片数据
func parseHouseImages(contents *[]byte, re *regexp.Regexp) []map[string]string {
	var matches = re.FindAllSubmatch(*contents, -1)
	var result = make([]map[string]string, 0)
	for _, match := range matches {
		if match == nil {
			continue
		}
		url := string(match[1])
		name := string(match[2])
		result = append(result, map[string]string{url: name})
	}
	return result
}

// parseHouseBasicInfo获取二手房详情页面基本信息数据
func parseHouseBasicInfo(contents *[]byte, title string) (*model.Ershoufang, error) {
	reader := bytes.NewReader(*contents)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	var base = make(map[string]string)
	var transaction = make(map[string]string)
	var baseTits = make([]string, 0)
	var baseVals = make([]string, 0)
	var transactionTits = make([]string, 0)
	var transactionVals = make([]string, 0)
	doc.Find(".baseinform div .introContent div").Find(".content").Each(func(index int, selection *goquery.Selection) {
		// 基本属性
		if index == 0 {
			selection.Find("li span").Each(func(si int, ss *goquery.Selection) {
				baseTits = append(baseTits, strings.TrimSpace(ss.Text()))
			})
			selection.Find("li").Each(func(li int, ls *goquery.Selection) {
				baseVals = append(baseVals, strings.TrimSpace(ls.Text()))
			})
		}

		// 交易属性
		if index == 1 {
			selection.Find("li span").Each(func(si int, ss *goquery.Selection) {
				if si%2 == 0 {
					transactionTits = append(transactionTits, strings.TrimSpace(ss.Text()))
				} else {
					transactionVals = append(transactionVals, strings.TrimSpace(ss.Text()))
				}
			})
		}
	})

	for index, title := range baseTits {
		base[title] = baseVals[index]
	}

	// marshal, err := json.Marshal(transactionTits)
	// fmt.Printf("transactionTits: %+v, len: %d, title: %s\n", string(marshal), len(transactionTits), title)
	// marshal, err = json.Marshal(transactionVals)
	// fmt.Printf("transactionVals: %+v, len: %d, title: %s\n", string(marshal), len(transactionVals), title)

	tlen := len(transactionTits)
	vlen := len(transactionVals)
	if tlen != vlen {
		if tlen > vlen {
			for index, value := range transactionVals {
				transaction[transactionTits[index]] = value
			}
		}
	} else {
		for index, title := range transactionTits {
			transaction[title] = transactionVals[index]
		}
	}

	// marshal, err = json.Marshal(transaction)
	// fmt.Printf("transaction: %+v, len: %d, title: %s\n", string(marshal), len(transaction), title)

	return &model.Ershoufang{
		BaseInfo: model.ErshoufangBaseInfo{
			Base:        base,
			Transaction: transaction,
		},
	}, nil
}

// parseHouseBasicInfo获取二手房详情页面房源特色数据
func parseHouseCharacteristics(contents *[]byte) []map[string]string {
	reader := bytes.NewReader(*contents)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		panic(err)
	}
	var characteristics = make([]map[string]string, 0)
	var titS = make([]string, 0)
	var valS = make([]string, 0)
	doc.Find(".box-l div .introContent.showbasemore div").Find("div").Each(func(index int, selection *goquery.Selection) {
		content := strings.TrimSpace(selection.Text())
		if index%2 == 0 {
			titS = append(titS, content)
		} else {
			if selection.Find("a").Nodes != nil {
				// 房源标签
				var tags = make([]string, 0)
				selection.Find("a").Each(func(ai int, as *goquery.Selection) {
					tags = append(tags, strings.TrimSpace(as.Text()))
				})
				if tags != nil {
					valS = append(valS, strings.Join(tags, "|"))
				}
			} else {
				valS = append(valS, content)
			}
		}
	})
	for index, tit := range titS {
		characteristics = append(characteristics, map[string]string{tit: valS[index]})
	}
	return characteristics
}

// parseHouseFields 获取二手房详情页面基本字段数据
func parseHouseFields(contents *[]byte, re *regexp.Regexp) string {
	matchs := re.FindSubmatch(*contents)
	// 如：碧桂园1号公园
	if len(matchs) != 0 && len(matchs) == 2 {
		return string(matchs[1])
	}
	// 如：迎江区-龙狮桥乡
	if len(matchs) != 0 && len(matchs) == 3 {
		name := string(matchs[1])
		if string(matchs[2]) != "" {
			name += "-" + string(matchs[2])
		}
		return name
	}
	return ""
}
