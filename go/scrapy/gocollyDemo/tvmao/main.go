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
	sample.StartCollectV2()
}
