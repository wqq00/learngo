package main

import (
	//"learngo/helloword/crawler/engine"
	//"learngo/helloword/crawler/zhenai/parser"
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
	"io"
	"bufio"
	"regexp"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error: status code", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil{
		panic(err)
	}
	ParseCityList(all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

//const cityList = `<a  href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>([^<]+</a>`

func ParseCityList(contents []byte)  {
	re := regexp.MustCompile(`(<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _, m := range matches {
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}

	fmt.Printf("Matches found: %d\n", len(matches))
}

