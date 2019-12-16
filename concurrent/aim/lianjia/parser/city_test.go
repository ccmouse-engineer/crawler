package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	c, err := ioutil.ReadFile("city_test_data.txt")
	if err != nil {
		panic(err)
	}
	const resultRequestSize = 129
	expectedUrls := []string{
		"https://aq.lianjia.com/ershoufang/103104788556.html",
		"https://aq.lianjia.com/ershoufang/103105258596.html",
		"https://aq.lianjia.com/ershoufang/103105442776.html",
	}
	parseResult := ParseCity(c, "安庆", "https://aq.lianjia.com/", true)
	if len(parseResult.Requests) != resultRequestSize {
		t.Errorf("result should have %d requests; but had %d\n", resultRequestSize, len(parseResult.Requests))
	}

	for idx, url := range expectedUrls {
		if parseResult.Requests[idx].Url != url {
			t.Errorf("expected url #%d: %s; but was %s\n", idx, url, parseResult.Requests[idx].Url)
		}
	}
}
