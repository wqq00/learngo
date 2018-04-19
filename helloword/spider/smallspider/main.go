//抓去页面里所有的图片并打包
package main

import (
	"path/filepath"
	"net/http"
	"log"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"net/url"
	"strings"
	"os"
	"io"
	"io/ioutil"
	"compress/gzip"
	"archive/tar"
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
	doc.Find("img").Each(func(i int, s *goquery.Selection){
		link, ok := s.Attr("src")
		if ok {
			urls = append(urls, link)
		} else{
			fmt.Println("no link")
		}
	})
	return urls, nil
}

func make_tar(dir string, w io.Writer) error{
	basedir := filepath.Base(dir) //读取文件目录
	compress := gzip.NewWriter(w) //实现压缩功能
	defer compress.Close()
	tw := tar.NewWriter(w) //表示我们会把数据都写入w中去。而这个w就是我们在主函数中创建的文件。
	defer tw.Close()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error{ //用"filepath.Walk"函数去递归"dir"目录下的文件
		header, err := tar.FileInfoHeader(info, "") //将文件信息读取出来传给"header"变量，注意"info"后面参数是不传值的，除非你的目录是该软连接。
		if err != nil{
			return err
		}
		p, _ := filepath.Rel(dir, path) //取出文件的上级目录。
		//header.Name = path
		// 这是将path的相对路径传给"header.Name ",然后在写入到tw中去。不然的话只能拿到"info.Name ()"的名字，也就是如果不来这个操作的话它只会保存文件名，而不会记得路径。
		header.Name = filepath.Join(basedir, p)
		tw.WriteHeader(header) //将文件的信息写入到文件w中去。
		if info.IsDir(){
			return nil
		}
		f1, err := os.Open(path)
		if err != nil{
			log.Panic("创建文件出错!")
		}
		defer f1.Close()
		io.Copy(tw,f1)
		return nil
	})

	return nil
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

func downloadImgs(urls []string, dir string) error{
	for _, link := range urls{
		resp, err := http.Get(link)
		if err != nil{
			continue
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK{
			log.Fatal(resp.Status)
		}
		file_name := filepath.Base(link)
		full_name := filepath.Join(dir, file_name)
		f, err := os.Create(full_name)
		if err != nil{
			log.Panic("create fail")
		}
		io.Copy(f, resp.Body)
		fmt.Printf("已下载文件至：\033[31;1m%s\033[0m\n",full_name)
	}
	return nil
}

func main(){
	root_path := "http://daily.zhihu.com/"
	picture_path, err := fetch(root_path)
	if err != nil{
		log.Fatal(err)
	}
	if err != nil{
		log.Fatal(err)
	}

	Absolute_path := Clean_urls(root_path, picture_path)
	tmpdir, err := ioutil.TempDir("E:\\gocode\\download", "")
	fmt.Printf("%v", Absolute_path)
	fmt.Println(tmpdir)
	err = downloadImgs(Absolute_path, tmpdir)
	if err != nil{
		log.Panic(err)
	}

	//make_tar("..", os.Stdout)  //将结果输出到屏幕上
	f, err := os.Create("E:\\gocode\\download\\img.tar.gz") //
	if err != nil{
		fmt.Println(err)
	}
	defer f.Close()
	make_tar(tmpdir, f)

}







