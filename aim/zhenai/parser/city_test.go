package parser

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestUsersMergeOfCity(t *testing.T) {
	c, err := ioutil.ReadFile("city_test_data_of_source.html")
	if err != nil {
		panic(err)
	}
	cities := UsersMergeOfCity(c)
	expected, err := ioutil.ReadFile("city_test_data_of_expected.json")
	var data interface{}
	err = json.Unmarshal(expected, &data)
	if err != nil {
		panic(err)
	}
	data, ok := data.([]interface{})
	if ok {
		for index, datum := range data.([]interface{}) {
			datum, ok = datum.(map[string]interface{})
			if ok {
				if cities[index].Name != datum.(map[string]interface{})["Name"] {
					t.Errorf("expected city of index %d Name field value %s, but has %s",
						index, datum.(map[string]interface{})["Name"], cities[index].Name)
				}

				if cities[index].URL != datum.(map[string]interface{})["URL"] {
					t.Errorf("expected city of index %d URL field value %s, but has %s",
						index, datum.(map[string]interface{})["URL"], cities[index].URL)
				}

				if cities[index].Gender != datum.(map[string]interface{})["Gender"] {
					t.Errorf("expected city of index %d Gender field value %s, but has %s",
						index, datum.(map[string]interface{})["Gender"], cities[index].Gender)
				}
			}
		}
	}
}
