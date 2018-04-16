package scheduler

import "learngo/helloword/crawler_2/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request)  {
	// send request down to worker chan
	s.workerChan <- r
}