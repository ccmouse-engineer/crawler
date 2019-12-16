package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCar(t *testing.T) {
	c, err := ioutil.ReadFile("car_test_data.html")
	if err != nil {
		panic(err)
	}

	const resultSize = 56
	expectedUrls := []string{
		"http://newcar.xcar.com.cn/car/select/s2547/",
		"http://newcar.xcar.com.cn/car/select/s2365/",
		"http://newcar.xcar.com.cn/car/select/s553/",
	}
	expectedCities := []string{
		"奥迪|奥迪A3三厢",
		"奥迪|奥迪A3两厢",
		"奥迪|奥迪A4L",
	}
	parseResult := ParseCar(c, domain, "奥迪")
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
