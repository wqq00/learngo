package main

import (
	"time"
	"learngo/learngo/helloword/oldboy/day9/chat/chat_server/maingo"
)


func main() {
	maingo.InitRedis("localhost:6379", 16, 1024, time.Second*300)
	maingo.InitUserMgr()
	maingo.RunServer("0.0.0.0:10000")
}

