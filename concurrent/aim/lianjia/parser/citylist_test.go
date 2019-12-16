package parser

import (
	"io/ioutil"
	"testing"
)

// TestParseCityList test case
func TestParseCityList(t *testing.T) {
	c, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		t.Fatalf("fetcher.Fetch error: %v\n", err)
	}
	const resultRequestSize = 105
	expectedUrls := []string{
		"https://aq.lianjia.com/ershoufang/",
		"https://baoji.lianjia.com/ershoufang/",
		"https://bd.lianjia.com/ershoufang/",
	}
	parseResult := ParseCityList(c, "ershoufang", -1)
	if len(parseResult.Requests) != resultRequestSize {
		t.Errorf("result should have %d requests; but had %d\n", resultRequestSize, len(parseResult.Requests))
	}
	for idx, url := range expectedUrls {
		if parseResult.Requests[idx].Url != url {
			t.Errorf("expected url #%d: %s; but was %s\n", idx, url, parseResult.Requests[idx].Url)
		}
	}
}
