package main

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func checkFileIsExist(filename string) bool{
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err){
		exist = false
	}
	return exist
}

func main(){
	var wireteString = "测试n"
	var filename = "./output1.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	check(err1)
	fmt.Printf("写入 %d 个字节n", n)

	var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile("./output2.txt", d1, 0666) //写入文件(字节数组)
	check(err2)

}