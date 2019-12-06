package engine

import "log"

// ConcurrentEngine并发引擎结构体类型
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler调度器接口类型
type Scheduler interface {
	Submit(Request)                         // 注册请求
	ConfigureMasterWorkerChan(chan Request) // 工作者数据处理通道
}

// Run爬虫入口
func (c *ConcurrentEngine) Run(seeds ...Request) {
	// 创建收发数据通道
	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureMasterWorkerChan(in)

	// 创建工作者处理请求
	for i := 0; i < c.WorkerCount; i++ {
		createdWorker(in, out)
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
func createdWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
