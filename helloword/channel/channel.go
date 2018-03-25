package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorkder(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int //c == nil
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorkder(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func Hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func bufferedChannel() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelCloser() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelCloser()
}
