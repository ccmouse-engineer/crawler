package parser

import (
	"io/ioutil"
	"testing"
)

const domain = "http://newcar.xcar.com.cn"

// TestParseCityList test case
func TestParseBrandList(t *testing.T) {
	c, err := ioutil.ReadFile("brandlist_test_data.html")
	if err != nil {
		t.Fatalf("fetcher.Fetch error: %v\n", err)
	}

	const resultSize = 190
	expectedUrls := []string{
		"http://newcar.xcar.com.cn/car/0-0-0-0-1-0-0-0-0-0-0-0/",
		"http://newcar.xcar.com.cn/car/0-0-0-0-56-0-0-0-0-0-0-0/",
		"http://newcar.xcar.com.cn/car/0-0-0-0-78-0-0-0-0-0-0-0/",
	}
	expectedCities := []string{
		"奥迪",
		"阿斯顿·马丁",
		"阿尔法·罗密欧",
	}
	parseResult := ParseBrandList(c, domain)
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
