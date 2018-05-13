package engine

import (
	"log"
	"learngo/helloword/crawler/model"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	//ConfigureMasterWorkerChan(chan Request)
	Run()
}

type ReadyNotifier interface{
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {
	//in := make(chan Request)
	out := make(chan  ParserResult)
	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds{
		if isDuplicate(r.Url){
			log.Printf("Duplicate request: "+"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items{
			if _, ok := item.(model.Profile); ok{
				log.Printf("Got profile #%d: %v", profileCount, item)
				profileCount++
			}

		}

		// URL defup

		for _, request := range result.Requests{
			if isDuplicate(request.Url){
				log.Printf("Duplicate request: "+"%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}

}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier){
	//in := make(chan Request)
	go func(){
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)
func isDuplicate(url string) bool{
	if visitedUrls[url]{
		return true
	}
	visitedUrls[url] = true
	return false
}