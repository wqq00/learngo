package main

import (
	"fmt"
	"sync"
)

func dowork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorkder(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go dowork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	//var c chan int //c == nil
	var workers [10]worker

	for i := 0; i < 10; i++ {
		workers[i] = createWorkder(i, &wg)
	}

	wg.Add(10)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}

}

func main() {
	chanDemo()
}
