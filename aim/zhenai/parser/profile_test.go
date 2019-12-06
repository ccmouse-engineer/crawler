package parser

import (
	"crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfileForInfo(t *testing.T) {
	// source code url: https://album.zhenai.com/u/1384450731
	c, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		t.Errorf("fetcher.Fetch error: %s\n", err)
	}
	expectedHTML := `<div class="purple-btns" data-v-8b1eac0c=""><div class="m-btn purple" data-v-8b1eac0c="">离异</div><div class="m-btn purple" data-v-8b1eac0c="">31岁</div><div class="m-btn purple" data-v-8b1eac0c="">天秤座(09.23-10.22)</div><div class="m-btn purple" data-v-8b1eac0c="">158cm</div><div class="m-btn purple" data-v-8b1eac0c="">45kg</div><div class="m-btn purple" data-v-8b1eac0c="">工作地:阿坝汶川</div><div class="m-btn purple" data-v-8b1eac0c="">月收入:3-5千</div><div class="m-btn purple" data-v-8b1eac0c="">自由职业</div><div class="m-btn purple" data-v-8b1eac0c="">大专</div></div> <div class="pink-btns" data-v-8b1eac0c=""><div class="m-btn pink" data-v-8b1eac0c="">藏族</div><div class="m-btn pink" data-v-8b1eac0c="">籍贯:四川阿坝</div><div class="m-btn pink" data-v-8b1eac0c="">体型:瘦长</div><div class="m-btn pink" data-v-8b1eac0c="">不吸烟</div><div class="m-btn pink" data-v-8b1eac0c="">稍微喝一点酒</div><div class="m-btn pink" data-v-8b1eac0c="">和家人同住</div><div class="m-btn pink" data-v-8b1eac0c="">未买车</div><div class="m-btn pink" data-v-8b1eac0c="">有孩子且住在一起</div><div class="m-btn pink" data-v-8b1eac0c="">是否想要孩子:视情况而定</div><div class="m-btn pink" data-v-8b1eac0c="">何时结婚:时机成熟就结婚</div></div>`
	info, err := ParseProfileForInfo(c)
	if err != nil {
		t.Errorf("parse user profile infomation error: %s\n", err)
	}
	if string(info) != expectedHTML {
		t.Errorf("parse user profile infomatin faild. got: %s\n", string(info))
	}
}

func TestParseProfile(t *testing.T) {
	// source code url: https://album.zhenai.com/u/1384450731
	c, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		t.Fatalf("fetcher.Fetch error: %s\n", err)
	}
	parseResult := ParseProfile(c, "墨洁", "女士")
	if len(parseResult.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v\n", parseResult.Items)
	}
	expected := model.Profile{
		Name:            "墨洁",
		Age:             31,
		Gender:          "女士",
		Height:          158,
		Weight:          45,
		WorkingLocation: "工作地:阿坝汶川",
		Income:          "月收入:3-5千",
		Marriage:        "离异",
		Education:       "大专",
		Occupation:      "自由职业",
		Hometown:        "籍贯:四川阿坝",
		Nation:          "藏族",
		Constellation:   "天秤座(09.23-10.22)",
		House:           "和家人同住",
		Car:             "未买车",
		Drinking:        "稍微喝一点酒",
		Smoking:         "不吸烟",
	}

	profile, ok := parseResult.Items[0].(model.Profile)
	if !ok {
		t.Errorf("incorrect item value type: %+v\n", parseResult.Items[0])
	}

	if expected != profile {
		t.Errorf("expected %v, but was %v\n", expected, profile)
	}

}
