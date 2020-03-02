package searchServer

import "gocollyDemo/tvmao/sample"

func StartSearch(key string) (ok bool){

	urlDrama := tvmao.Search(key)
	savePathDir := "C:/ZTestData/TVMao/"
	tvmao.StartCollect(urlDrama,savePathDir)

	ok = true
	return
}
