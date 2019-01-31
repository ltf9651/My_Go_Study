package scheduler

import (
	"My_Go_Study/Crawler/engine"
)

type SimpleScheduler struct {
	workerChannel chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker channel
	// s.workerChannel <- r 出现卡死
	go func() {
		s.workerChannel <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChannel
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChannel = make(chan engine.Request)
}
