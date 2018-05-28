package main

import (
	"flag"
	"time"
	"github.com/sirupsen/logrus"
	"os"
	"bufio"
	"io"
	"strings"
	"net/url"
	"crypto/md5"
	"encoding/hex"
)

const HANDLE_DIG = "/dig?"

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
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main(){

	//获取参数
	logFilePath := flag.String("logFilePah", "/User/pangee/Public/nginx/log/dig.log", "log file path")
	routineNum := flag.Int("routineNum", 5, "consumer numble by goroutine")
	l := flag.String("l", "/temp/log", "this programe runtime log target file path")
	flag.Parse()

	params := cmdParams{*logFilePath, *routineNum}

	//打日志
	logFd, err := os.OpenFile(*l, os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil{
		log.Out = logFd
		defer logFd.Close()
	}

	log.Infoln("Exec start.")
	log.Infoln("Param:logFilePath=%s, routineNum=%d", params.logFilePath, params.routineNum)

	//初始化一些channel，用于数据传递
	var logChannel = make(chan string, 3*params.routineNum)
	var pvChannel = make(chan urlData, params.routineNum)
	var uvChannel = make(chan urlData, params.routineNum)
	var storageChannel = make(chan storageBlock, params.routineNum)

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

func uvCounter(uvChnanel chan urlData, storageChannel chan storageBlock){

}

func logConsumer(logChannel chan string, pvChannel, uvChnanel chan urlData) error {
	for logStr := range logChannel{
		// 切割日志字符串，抠出打点上报的日志
		data := cutLogFetchData(logStr)

		//uid
		hasher := md5.New()
		hasher.Write([]byte(data.refer+data.ua))
		uid := hex.EncodeToString(hasher.Sum(nil))

		uData := urlData{data, uid}
		pvChannel <- uData
		uvChnanel <- uData
	}
	return nil
}

func cutLogFetchData(logStr string) digData{
	logStr = strings.TrimSpace(logStr)
	pos1 := str.IndexOf(logStr, HANDLE_DIG, 0)
	if pos1 == -1{
		return digData{}
	}
	pos1 += len(HANDLE_DIG)
	pos2 := str.IndexOf(logStr, "HTTP/", pos1)
	d := str.Substr(logStr, pos1, pos2-pos1)

	urlInfo, err := url.Parse("http://localhost/?"+d)
	if err != nil{
		return digData{}
	}
	data := urlInfo.Query()
	return digData{
		data.Get("time"),
		data.Get("refer"),
		data.Get("url"),
		data.Get("ua"),
	}
}

func readFileLinenyLine(params cmdParams, logChannal chan string) error {
	fd, err := os.Open(params.logFilePath)
	if err != nil{
		log.Warning("ReadFileLinebyLine can't open file:%s", params.logFilePath)
		return err
	}
	defer fd.Close()

	bufferRead := bufio.NewReader(fd)

	count := 0
	for {
		line, err := bufferRead.ReadString('\n')
		logChannal <- line
		count++
		if count%(1000*params.routineNum) == 0{
			log.Infof("ReadFileLinebyLine line:%d", count)
		}
		if err != nil{
			if err == io.EOF{
				time.Sleep(3*time.Second)
				log.Infof("ReadFileLinebyLine wait, readline:%d", count)
			}else{
				log.Warning("ReadFileLinebyLine read log err")
			}
		}
	}
	return nil
}