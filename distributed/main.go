package main

import (
	"crawler/concurrent/aim/lianjia/parser"
	"crawler/concurrent/engine"
	"crawler/concurrent/scheduler"
	"crawler/distributed/config"
	"crawler/distributed/persist/client"
	"fmt"
)

func main() {
	// 获取存储爬取项服务
	itemSaver, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
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
