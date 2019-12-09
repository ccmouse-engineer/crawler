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
	const resultSize = 105
	expectedUrls := []string{
		"https://aq.lianjia.com/ershoufang/",
		"https://baoji.lianjia.com/ershoufang/",
		"https://bd.lianjia.com/ershoufang/",
	}
	expectedCities := []string{
		"安庆",
		"宝鸡",
		"保定",
	}
	parseResult := ParseCityList(c, "ershoufang", -1)
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
