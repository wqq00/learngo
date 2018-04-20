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

	root_path := "http://www.w3school.com.cn/tags/tag_tr.asp"
	fetch(root_path)
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


}
