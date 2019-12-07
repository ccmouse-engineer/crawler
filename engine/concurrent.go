package engine

import (
	"log"
)

// ConcurrentEngine并发引擎结构体类型
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler调度器接口类型
type Scheduler interface {
	Submit(Request)                         // 注册请求
	ConfigureMasterWorkerChan(chan Request) // 工作者数据处理主通道
	WorkerReady(chan Request)
	Run()
}

// Run爬虫入口
func (c *ConcurrentEngine) Run(seeds ...Request) {
	// 创建收数据通道
	out := make(chan ParseResult)
	c.Scheduler.Run()

	// 创建工作者处理请求
	for i := 0; i < c.WorkerCount; i++ {
		createdWorker(out, c.Scheduler)
	}

	// 将请求种子放入到调度器
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	// 接收请求处理后数据
	itemNumber := 1
	for {
		parseResult := <-out
		// 展示处理结果
		for _, item := range parseResult.Items {
			log.Printf("Got item #%d: %+v", itemNumber, item)
			itemNumber++
		}

		// 注册处理结果中的请求到调度器
		for _, r := range parseResult.Requests {
			c.Scheduler.Submit(r)
		}
	}

}

// createdWorker创建工作者处理请求
func createdWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
