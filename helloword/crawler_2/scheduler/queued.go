package scheduler

import (
	"learngo/helloword/crawler_2/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request){
	s.workerChan <- w
}

//func (s *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
//	panic("implement me")
//}

func (s *QueuedScheduler) Run(){
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan  engine.Request)
	go func(){
		var requestQ []engine.Request
		var workderQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workderQ) > 0 {
				activeWorker = workderQ[0]
				activeRequest = requestQ[0]
			}
			select{
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workderQ = append(workderQ, w)
			case activeWorker <- activeRequest:
				workderQ = workderQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}
