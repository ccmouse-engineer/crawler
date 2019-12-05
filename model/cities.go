package model

// City城市下每一页面每一个用户
type City struct {
	Name   string
	URL    string
	Gender string
}

// Cities城市下每一页面所有用户
type Cities []City
