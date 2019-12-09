package engine

import (
	"crawler/model"
	"encoding/json"
	"log"
)

// accessedURLs访问过的URL
var accessedURLs = make(map[string]bool)

// ConcurrentEngine并发引擎结构体类型
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

// Scheduler调度器接口类型
type Scheduler interface {
	ReadyNotifier
	Submit(Request) // 注册请求
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run爬虫入口
func (c *ConcurrentEngine) Run(seeds ...Request) {
	// 创建收数据通道
	out := make(chan ParseResult)
	c.Scheduler.Run()

	// 创建工作者处理请求
	for i := 0; i < c.WorkerCount; i++ {
		createdWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	// 将请求种子放入到调度器
	for _, r := range seeds {
		if isDuplicate(r) {
			continue
		}
		c.Scheduler.Submit(r)
	}

	// 接收请求处理后数据
	houseNum := 1
	for {
		parseResult := <-out
		// 展示处理结果
		for _, item := range parseResult.Items {
			if ershoufang, ok := item.(*model.Ershoufang); ok {
				marshal, _ := json.Marshal(ershoufang)
				log.Printf("Got item #%d: %+v\n", houseNum, string(marshal))
				houseNum++
			}
		}

		// 注册处理结果中的请求到调度器
		for _, r := range parseResult.Requests {
			if isDuplicate(r) {
				continue
			}
			c.Scheduler.Submit(r)
		}
	}

}

// 检测是不是重复的URL
func isDuplicate(r Request) bool {
	if accessedURLs[r.Url] {
		return true
	}
	accessedURLs[r.Url] = true
	return false
}

// createdWorker创建工作者处理请求
func createdWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
