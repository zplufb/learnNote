package main

import (
			"Test001/golang/scrapy/gocollyDemo/tvmao/sample"
)

func main() {
	//StartCollect()
	//搜索获取其相关信息
	key := "九州缥缈录"
	urlDrama := tvmao.Search(key)
	savePathDir := "E:/ZTestData/TVMao/"
	tvmao.StartCollect(urlDrama,savePathDir)
}
