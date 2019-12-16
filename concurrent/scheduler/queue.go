package scheduler

import (
	"crawler/concurrent/engine"
)

// QueueScheduler调度器接口的队列实现类型
type QueueScheduler struct {
	requestChan chan engine.Request      // 请求队列通道
	workerChan  chan chan engine.Request // 工作队列通道
}

// WorkerChan提供一个新的工作者通道(传递的是请求)
func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// Submit注册请求到请求队列
func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

// WorkerReady指明一个工作者准备就绪可以开始工作了
func (q *QueueScheduler) WorkerReady(worker chan engine.Request) {
	q.workerChan <- worker
}

// Run运行调度器
func (q *QueueScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-q.requestChan:
				// 放入到请求队列
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				// 放入到工作者队列
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				// 发送之后从队列中取出
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
