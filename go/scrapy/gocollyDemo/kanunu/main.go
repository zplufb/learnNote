package main

import "Test001/golang/scrapy/gocollyDemo/kanunu/smaple"

func main() {
	//收集小说每章节的链接
	var UrlIndex = "https://www.kanunu8.com/book2/11082/index.html"
	kanunu.CollectLinks(UrlIndex)
	//kanunu.ReadNovelDemo1()
}
