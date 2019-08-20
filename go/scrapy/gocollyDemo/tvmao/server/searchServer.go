package searchServer

import "Test001/golang/scrapy/gocollyDemo/tvmao/sample"

func StartSearch(key string) (ok bool){

	urlDrama := tvmao.Search(key)
	savePathDir := "E:/ZTestData/TVMao/"
	tvmao.StartCollect(urlDrama,savePathDir)

	ok = true
	return
}
