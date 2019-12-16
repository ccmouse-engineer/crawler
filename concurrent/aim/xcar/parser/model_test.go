package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseModel(t *testing.T) {
	c, err := ioutil.ReadFile("model_test_data.html")
	if err != nil {
		panic(err)
	}
	const resultSize = 18
	expectedUrls := []string{
		"http://newcar.xcar.com.cn/m51092/",
		"http://newcar.xcar.com.cn/m51093/",
		"http://newcar.xcar.com.cn/m51094/",
	}
	expectedCities := []string{
		"2020款 Limousine 35 TFSI 进取型",
		"2020款 Limousine 35 TFSI时尚型",
		"2020款 Limousine 35 TFSI风尚型",
	}
	parseResult := ParseModel(c, domain, "奥迪A3三厢")
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
