package helper

import (
	"io/ioutil"
	"testing"
)

func TestParseOccupationToDelimitedStr(t *testing.T) {
	source, err := ioutil.ReadFile("profile_test_data_of_occupation.json")
	if err != nil {
		t.Error(err)
	}
	parseOccupationToDelimitedStr, err := ParseOccupationToDelimitedStr(string(source))
	if parseOccupationToDelimitedStr == "" {
		t.Errorf("result should have %s occupation; but had %s\n", parseOccupationToDelimitedStr, parseOccupationToDelimitedStr)
	}
}
