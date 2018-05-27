package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic("啊哦,出错了")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic("啊哦,又出错了")
	}
	fmt.Printf("%s/n", all)

}
