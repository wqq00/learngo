// 协程池

package main

import (
	"sync"
	"net/http"
	"log"
)

func work(ch chan string, wg *sync.WaitGroup){ //ch是用来接受数据，wg是用来同步协程的.
	for u := range ch{
		resp, err := http.Get(u)
		if err != nil{
			log.Print(err)
			return
		}
		log.Printf("访问的网站是：%s，网页的大小是:%v",u,resp.ContentLength)
		resp.Body.Close()
	}
	wg.Done()
}

func main(){
	wg := new(sync.WaitGroup)
	wg.Add(5)
	taskch := make(chan string)
	for i := 0; i < 5; i++{
		go work(taskch, wg)
	}
	urls := []string{"http://www.baidu.com","http://www.zhihu.com","http://www.cnblogs.com/yinzhengjie/p/7201980.html"}
	for _, url := range urls{
		taskch <- url
	}
	close(taskch)
	wg.Wait()

}
