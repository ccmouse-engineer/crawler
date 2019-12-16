package parser

import (
	"bytes"
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"errors"
	"fmt"
	"regexp"

	"github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
)

var (
	vendorGuidePriceRe = regexp.MustCompile(`<a\s+href="/\w+/baojia/"\s+target="_blank"\s+onclick="[^"]+">([^>]+[^<]+)</a>`)
	imageRe            = regexp.MustCompile(`<img\s+class="color_car_img_new"\s+src="(//[\w+./-]+)"`)
)

// ParseDetail解析HTTP响应内容汽车详情页
func ParseDetail(contents []byte, name string, model string) (parseResult engine.ParseResult) {
	// 解析车型页面厂商指导价
	imageURL, err := ParseCarImageURL(&contents, imageRe)
	if err != nil {
		logrus.Errorln(imageURL)
		logrus.Errorln(err)
	}

	// 解析车型页面厂商指导价
	vendorGuidePrice, err := ParseCarVendorGuidePrice(&contents, vendorGuidePriceRe)
	if err != nil {
		logrus.Errorln(err)
	}

	// 获取汽车详情页面参数(基本参数和车身参数)
	car, err := ParseParameterConfiguration(&contents)
	if err != nil {
		logrus.Errorln(err)
	}

	car.Name = model + "|" + name
	car.VendorGuidePrice = vendorGuidePrice
	car.ImageURL = imageURL
	parseResult.Items = append(parseResult.Items, car)
	return parseResult
}

// ParseImageURL 解析车型页面汽车图片
func ParseCarImageURL(contents *[]byte, r *regexp.Regexp) (string, error) {
	matches := r.FindSubmatch(*contents)
	if len(matches) > 1 {
		return "http:" + string(matches[1]), nil
	}
	return "", fmt.Errorf("match image url fatal error, please check regexp")
}

// ParseVendorGuidePrice 解析车型页面厂商指导价
func ParseCarVendorGuidePrice(contents *[]byte, r *regexp.Regexp) (string, error) {
	matches := r.FindSubmatch(*contents)
	if len(matches) > 1 {
		return string(matches[1]) + "万", nil
	}
	return "", errors.New("match vendor guide price fatal error, please check regexp")
}

// ParseParameterConfiguration获取汽车详情页面参数(基本参数和车身参数)
func ParseParameterConfiguration(contents *[]byte) (car *model.Car, err error) {
	reader := bytes.NewReader(*contents)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	baseInfo := make(map[string]string, 0)
	baseTitle := make([]string, 0)
	baseValue := make([]string, 0)
	bodyInfo := make(map[string]string, 0)
	bodyTitle := make([]string, 0)
	bodyValue := make([]string, 0)
	doc.Find("#div_newd0_js_2044").Next().Find("table").Each(func(index int, selection *goquery.Selection) {
		// 基本参数
		if index == 0 {
			selection.Find("tbody tr").Each(func(tri int, trs *goquery.Selection) {
				if tri > 0 {
					trs.Children().Each(func(tdi int, tds *goquery.Selection) {
						ret, _ := tds.Html()
						if tdi%2 == 0 {
							baseTitle = append(baseTitle, ret)
						} else {
							baseValue = append(baseValue, ret)
						}
					})
				}
			})
		}

		// 车身参数
		if index == 1 {
			selection.Find("tbody tr").Each(func(tri int, trs *goquery.Selection) {
				if tri > 0 {
					trs.Children().Each(func(tdi int, tds *goquery.Selection) {
						ret, _ := tds.Html()
						if tdi%2 == 0 {
							bodyTitle = append(bodyTitle, ret)
						} else {
							bodyValue = append(bodyValue, ret)
						}
					})
				}
			})
		}
	})

	for index, title := range baseTitle {
		baseInfo[title] = baseValue[index]
	}

	for index, title := range bodyTitle {
		bodyInfo[title] = bodyValue[index]
	}
	return &model.Car{Base: baseInfo, Body: bodyInfo}, nil
}
