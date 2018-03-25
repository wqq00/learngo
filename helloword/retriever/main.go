package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://wwww.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Println(download(r))
	r = real.Retriever{}
	fmt.Println(download(r))
}
