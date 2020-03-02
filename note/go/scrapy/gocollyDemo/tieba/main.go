package main

import "Test001/golang/scrapy/gocollyDemo/tieba/sample"

func main() {

	url := "https://tieba.baidu.com/p/4803144798"
	pageNum := tieba.GetTiebaPage(url)
	needNum := 10
	tieba.GetTiebaImages(url, pageNum, needNum)
}
