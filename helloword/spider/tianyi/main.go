package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/url"
	"strings"
	"path/filepath"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"log"
	"errors"
	"io/ioutil"
)

func fetch(url string) ([]string, error){

	var urls []string
	resp, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil, errors.New(resp.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil{
		log.Fatal(err)
	}

	text,_ := doc.Find(".dataintable").Find("tbody tr th").Eq(1).Html()
	fmt.Println(text)
	doc.Find(".ssc").Each(func(i int, s *goquery.Selection){
		link:= s.Find(".issue").Text()
		fmt.Println(link)


		//if ok {
		//	urls = append(urls, link)
		//} else{
		//	fmt.Println("no link")
		//}
	})
	return urls, nil
}

func Clean_urls(root_path string, picture_path []string) []string{
	var Absolute_path []string
	url_info, err := url.Parse(root_path)
	if err != nil{
		log.Fatal(err)
	}
	Scheme := url_info.Scheme
	Host := url_info.Host
	for _, souce_path := range picture_path{
		if strings.HasPrefix(souce_path, "https"){

		} else if strings.HasPrefix(souce_path, "//"){
			souce_path = Scheme + ":" + souce_path
		} else if strings.HasPrefix(souce_path, "/"){
			souce_path = Scheme + "//" + Host + souce_path
		} else {
			souce_path = filepath.Dir(root_path) + souce_path
		}
		Absolute_path = append(Absolute_path, souce_path)
	}
	return Absolute_path
}

var db = &sql.DB{}

func init(){
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tianyi?charset=utf8")
}

func main(){

	//root_path := "http://www.w3school.com.cn/tags/tag_tr.asp"
	//fetch(root_path)
	//picture_path, err := fetch(root_path)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//if err != nil{
	//	log.Fatal(err)
	//}


	//rows, _ := db.Query("select id, number, w, q  from test")
	//for rows.Next() {
	//	var id int
	//	var number int
	//	var w int
	//	var q int
	//	rows.Scan(&id, &number, &w, &q)
	//	fmt.Println(id)
	//	fmt.Println(number)
	//	fmt.Println(w)
	//	fmt.Println(q)
	//}

	//resp, err := http.Get("http://em.ttyule6.com/api/web-ajax/is-login")
	//if err != nil {
	//	// handle error
	//}
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//header := resp.Header
	////fmt.Printf("%#v",header["Set-Cookie"])
	//
	////re := regexp.MustCompile(header["Set-Cookie"])
	//fmt.Println(Between(header["Set-Cookie"][0], "SESSION=", "; Path=/; HttpOnly"))
	//fmt.Println(Between(header["Set-Cookie"][1], "__jsluid=", "; max-age=31536000"))
	//
	//cookiestrs := "__jsluid="+Between(header["Set-Cookie"][0], "SESSION=", "; Path=/; HttpOnly")+"; SESSION="+Between(header["Set-Cookie"][1], "__jsluid=", "; max-age=31536000")
	//fmt.Println(cookiestrs)

	client := &http.Client{}
	//post请求
	postValues := url.Values{}
	//curl 'http://ttyl.109705.com/api/web-login'
	// -H 'Cookie: __jsluid=70204eee8c4fee5e2d3fccf12cf8b5f6; SESSION=2405865d-90db-4106-9f70-a11ad6b025ac'
	// -H 'Origin: http://ttyl.109705.com'
	// -H 'Accept-Encoding: gzip, deflate'
	// -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8'
	// -H 'User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36'
	// -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8'
	// -H 'Accept: application/json, text/javascript, */*; q=0.01'
	// -H 'Referer: http://ttyl.109705.com/login.html'
	// -H 'X-Requested-With: XMLHttpRequest'
	// -H 'Connection: keep-alive'
	// --data 'username=wqq9898&password=qwer1234&securityCode=2902' --compressed

	//postValues.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	////postValues.Add("Cookie", cookiestrs)
	////postValues.Add("Cookie", "__jsluid=70204eee8c4fee5e2d3fccf12cf8b5f6; SESSION=2405865d-90db-4106-9f70-a11ad6b025ac")
	//postValues.Add("Host", "ttyl.109705.com")
	//postValues.Add("Origin", "http://ttyl.109705.com")
	//postValues.Add("Referer", "http://ttyl.109705.com/login.html")
	//postValues.Add("username", "wqq9898")
	//postValues.Add("password", "qwer1234")
	//postValues.Add("securityCode", "2654")
	//
	//resp2, err := client.PostForm("http://ttyl.109705.com/api/web-login", postValues)
	//defer resp2.Body.Close()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if resp2.StatusCode == 200 {
	//	body, _ := ioutil.ReadAll(resp2.Body)
	//	fmt.Println(string(body))
	//}

	//curl 'http://ttyl.109705.com/api/game-lottery/query-trend'
	// -H 'Cookie: __jsluid=70204eee8c4fee5e2d3fccf12cf8b5f6; SESSION=6e691e43-558d-4918-97bc-a572ec0e68c5'
	// -H 'Origin: http://ttyl.109705.com' -H 'Accept-Encoding: gzip, deflate'
	// -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8'
	// -H 'User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36'
	// -H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' -H 'Accept: application/json, text/javascript, */*; q=0.01'
	// -H 'Referer: http://ttyl.109705.com/game/lottery/trend.html?qqmin' -H 'X-Requested-With: XMLHttpRequest'
	// -H 'Connection: keep-alive'
	// --data 'name=qqmin&query=latest-30' --compressed

	postValues.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	postValues.Add("Accept-Encoding", "gzip, deflate")
	postValues.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	postValues.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.119 Safari/537.36")
	postValues.Add("Connection", "keep-alive")
	postValues.Add("Content-Length", "26")
	postValues.Add("X-Requested-With", "XMLHttpRequest")

	postValues.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//postValues.Add("Cookie", cookiestrs)
	postValues.Add("Cookie", "__jsluid=70204eee8c4fee5e2d3fccf12cf8b5f6; SESSION=6e691e43-558d-4918-97bc-a572ec0e68c5")
	postValues.Add("Host", "ttyl.109705.com")
	postValues.Add("Origin", "http://ttyl.109705.com")
	postValues.Add("Referer", "http://ttyl.109705.com/game/lottery/trend.html?qqmin")
	postValues.Add("name", "qqmin")
	postValues.Add("query", "latest-30")

	resp2, err := client.PostForm("http://ttyl.109705.com/api/game-lottery/query-trend", postValues)
	defer resp2.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp2.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp2.Body)
		fmt.Println(string(body))
	}


}

func Between(str, starting, ending string) string {
	s := strings.Index(str, starting)
	if s < 0 {
		return ""
	}
	s += len(starting)
	e := strings.Index(str[s:], ending)
	if e < 0 {
		return ""
	}
	return str[s : s+e]
}
