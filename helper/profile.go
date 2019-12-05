package helper

import (
	"encoding/json"
	"fmt"
)

// ParseOccupationToDelimitedStr 解析职业类型JSON数据到字符串
func ParseOccupationToDelimitedStr(data string) (string, error) {
	var d = []byte(data)
	var parse interface{}
	err := json.Unmarshal(d, &parse)
	if err != nil {
		return "", err
	}
	sum := 0
	result := ""
	items, ok := parse.(map[string]interface{})
	if ok {
		for _, item := range items {
			item, ok := item.([]interface{})
			if ok {
				for _, one := range item {
					one, ok := one.(map[string]interface{})
					if ok {
						if one["name"] == "其他" {
							break
						}
						name := ""
						if sum == 0 {
							name = fmt.Sprintf("%s", one["name"])
						} else {
							name = fmt.Sprintf("|%s", one["name"])
						}
						result += name
						sum++
					} else {
						return "", fmt.Errorf("parse data type error")
					}
				}
			} else {
				return "", fmt.Errorf("parse data type error")
			}
		}
	} else {
		return "", fmt.Errorf("parse data type error")
	}
	return result, nil
}
