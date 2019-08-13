package main

import (
	"os"
	"fmt"
	"Test001/golang/scrapy/gocollyDemo/tvmao/sample"
)

var fileName = "E:/ZTestData/tvmao_sjdzd2.txt"

func init() {
	//删除指定文件
	err := os.Remove(fileName)
	if err != nil {
		fmt.Errorf("err=%v", err)
	}

}
func main() {
	//StartCollect()
	//搜索获取其相关信息
	key := "天才"
	urlDrama := tvmao.Search(key)
	savePathDir := "E:/ZTestData/TVMao/"
	tvmao.StartCollect(urlDrama,savePathDir)
}
