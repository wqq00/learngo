package main

import (
	"flag"
	"time"
	"github.com/sirupsen/logrus"
)

type cmdParams struct {
	logFilePath string
	routineNum int
}

type digData struct {
	time string
	url string
	refer string
	ua string
}

type urlData struct {
	data digData
	uid string
}

type urlNode struct {

}

type storageBlock struct {
	counterType string
	storageModel string
	unode urlNode
}

var log = logrus.New()
func init(){

}

func main(){

	//获取参数
	logFilePath := flag.String("logFilePah", "/User/pangee/Public/nginx/log/dig.log", "log file path")
	routineNum := flag.Int("routineNum", 5, "consumer numble by goroutine")
	l := flag.String("l", "/temp/log", "this programe runtime log target file path")
	flag.Parse()

	params := cmdParams{*logFilePath, *routineNum}

	//打日志

	//初始化一些channel，用于数据传递
	var logChannel = make(chan string, 3*routineNum)
	var pvChannel = make(chan urlData, routineNum)
	var uvChannel = make(chan urlData, routineNum)
	var storageChannel = make(chan storageBlock, routineNum)

	//日志消费者
	go readFileLinenyLine(params, logChannel)

	//创建一组日志处理
	for i:=0; i<params.routineNum; i++{
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	//创建pv uv 统计器
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)
	//可扩展的

	//创建存储器
	go dataStorage(storageChannel)
	time.Sleep(1000*time.Second)

}

func dataStorage(storageChannel chan storageBlock){

}

func pvCounter(pvChnanel chan urlData, storageChannel chan storageBlock){

}

func uvCounter(pvChnanel chan urlData, storageChannel chan storageBlock){

}

func logConsumer(logChannel chan string, pvChannel, uvChnanel chan urlData){

}

func readFileLinenyLine(params cmdParams, logChannal chan string){

}