package main

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
	"net/url"
	"math/rand"
	"time"
	"os"
)

var uaList = []string{

}

type resource struct {
	url		string
	target	string
	start 	int
	end		int
}

func ruleResource() []resource{
	var res []resource
	r1 := resource{
		url : "http://muke_tongji.com/dig?aaa=as",
		target : "",
		start : 0,
		end : 0,
	}
	r2 := resource{
		url : "http://muke_tongji.com/dig?aaa={$id}",
		target : "{$id}",
		start : 1,
		end : 21,
	}
	r3 := resource{
		url : "http://muke_tongji.com/dig?bbb={$id}",
		target : "{$id}",
		start : 1,
		end : 12924,
	}
	res = append(res, r1)
	res = append(res, r2)
	res = append(res, r3)
	return res
}

func buildUrl(res []resource) []string{
	var list []string

	for _, r := range res{
		if len(r.target)==0{
			list = append(list, r.url)
		}else{
			for i:=r.start; i<=r.end; i++{
				urlStr := strings.Replace(r.url, r.target, strconv.Itoa(i), -1)
				list = append(list, urlStr)
			}
		}
	}
	return list
}

func makeLog(current, refer, ua string) string{
	u := url.Values{}
	u.Set("time", "1")
	u.Set("url", current)
	u.Set("refer", refer)
	u.Set("ua", ua)
	paramStr := u.Encode()

	logTemplate := ""
	log := strings.Replace(logTemplate, "{$paramStr}", paramStr, -1)
	log = strings.Replace(log, "{$ua}", ua, -1)
	return log
}

func randInt(min, max int) int{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if min > max{
		return max
	}
	return r.Intn(max-min) + min
}

func main(){
	total := flag.Int("total", 100, "how many rows by create")
	filePath := flag.String("filePath", "/var/log/nginx/access.log", "log file path")
	flag.Parse()

	res := ruleResource()
	list := buildUrl(res)

	logStr := ""
	for i:=0; i<=*total; i++{
		currentUrl := list[randInt(0, len(list)-1)]
		referUrl := list[randInt(0, len(list)-1)]
		ua := uaList[randInt(0, len(uaList)-1)]

		logStr = logStr + makeLog(currentUrl, referUrl, ua) + "\n"
		//ioutil.WriteFile(*filePath, []byte(logStr), 0644)
	}
	fd, _ := os.OpenFile( *filePath, os.O_RDWR|os.O_APPEND, 0644)
	defer fd.Close()
	fd.Write([]byte(logStr))


	fmt.Println("done.\n")
}
