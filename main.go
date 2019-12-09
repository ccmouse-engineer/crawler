package main

import (
	"crawler/aim/lianjia/parser"
	"crawler/engine"
	"crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
	}

	// 爱卡汽车
	//e.Run(engine.Request{
	//	Url: "http://newcar.xcar.com.cn/car/",
	//	ParserFunc: func(c []byte) engine.ParseResult {
	//		return parser.ParseBrandList(c, "http://newcar.xcar.com.cn")
	//	},
	//})

	// 链家二手房
	e.Run(engine.Request{
		Url: "https://www.lianjia.com/city/",
		ParserFunc: func(c []byte) engine.ParseResult {
			return parser.ParseCityList(c, "ershoufang", 1)
		},
	})
}
