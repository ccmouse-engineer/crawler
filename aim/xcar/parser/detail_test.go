package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseDetail(t *testing.T) {
	c, err := ioutil.ReadFile("detail_test_data.html")
	if err != nil {
		t.Errorf("fetcher.Fetch error: %s\n", err)
	}
	parseResult := ParseDetail(c, "2020款Limousine 35 TFSI 进取型", "13.66万起")
	fmt.Printf("%+v\n", parseResult)
}

func TestParseImageURL(t *testing.T) {
	c, err := ioutil.ReadFile("detail_test_data.html")
	if err != nil {
		t.Errorf("fetcher.Fetch error: %s\n", err)
	}
	expected := "http://img1.xcarimg.com/PicLib/s/s12468_300.jpg"
	s, err := ParseCarImageURL(&c, imageRe)
	if s != expected {
		t.Errorf("result should have %s; but had %s\n", expected, s)
	}
}

func TestParseVendorGuidePrice(t *testing.T) {
	c, err := ioutil.ReadFile("detail_test_data.html")
	if err != nil {
		t.Errorf("fetcher.Fetch error: %s\n", err)
	}
	expected := "19.52万"
	s, err := ParseCarVendorGuidePrice(&c, vendorGuidePriceRe)
	if s != expected {
		t.Errorf("result should have %s; but had %s\n", expected, s)
	}
}
