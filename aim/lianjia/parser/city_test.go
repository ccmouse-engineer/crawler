package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	c, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}
	const resultSize = 30
	expectedUrls := []string{
		"https://aq.lianjia.com/ershoufang/103104788556.html",
		"https://aq.lianjia.com/ershoufang/103105258596.html",
		"https://aq.lianjia.com/ershoufang/103105442776.html",
	}
	expectedCities := []string{
		"高区3房2厅景观房，可以一览长江全景",
		"碧桂园山水云间 3室2厅 102万",
		"碧桂园山水云间 3室2厅 102万",
	}
	parseResult := ParseCity(c, "安庆", "https://aq.lianjia.com/", true)
	if len(parseResult.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d\n", resultSize, len(parseResult.Items))
	}

	if len(parseResult.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d\n", resultSize, len(parseResult.Requests))
	}

	for idx, url := range expectedUrls {
		if parseResult.Requests[idx].Url != url {
			t.Errorf("expected url #%d: %s; but was %s\n", idx, url, parseResult.Requests[idx].Url)
		}
	}

	for idx, city := range expectedCities {
		if parseResult.Items[idx].(string) != city {
			t.Errorf("expected city #%d: %s; but was %s\n", idx, city, parseResult.Items[idx].(string))
		}
	}
}
