package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"bufio"
	"go-crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",

	})

	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic("啊哦,出错了")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	// 局限性,在不知道爬取网站编码时,无法正确执行
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	// 修改后
	e := determineEncoding(resp.Body)
	determineReader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(determineReader)
	if err != nil {
		panic("啊哦,又出错了")
	}
	fmt.Printf("%s/n", all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
