package scheduler

import (
	"crawler/engine"
)

// QueueScheduler调度器接口的简单实现类型
type SimpleScheduler struct {
	workerChan chan engine.Request
}

// WorkerChan提供一个新的工作者通道(简单实现无需实现该接口方法)
func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

// WorkerChan提供一个新的工作者通道(传递的是请求)
func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

// Run运行调度器
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// Submit注册请求到通道
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}
