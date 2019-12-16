package main

import (
	"crawler/concurrent/aim/lianjia/parser"
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"
	"crawler/concurrent/scheduler"
)

func main() {
	// 获取存储爬取项服务
	itemSaver, err := persist.ItemSaver("lianjia_test")
	if err != nil {
		panic(err)
	}

	// 开启并发爬取数据引擎
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemSaver,
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
