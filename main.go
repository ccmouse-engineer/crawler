package main

import (
	"crawler/aim/xcar/parser"
	"crawler/engine"
	"crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url: "http://newcar.xcar.com.cn/car/",
		ParserFunc: func(c []byte) engine.ParseResult {
			return parser.ParseBrandList(c, "http://newcar.xcar.com.cn")
		},
	})
}
