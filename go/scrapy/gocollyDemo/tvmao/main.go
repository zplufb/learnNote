package main

import (
	"net/http"
	"fmt"
	"Test001/golang/scrapy/gocollyDemo/tvmao/sample"
	"Test001/golang/scrapy/gocollyDemo/tvmao/controller"
)

func init() {
	http.HandleFunc("/query", searchController.Search)

	h := http.FileServer(http.Dir("/Go-Work/src/Test001/golang/scrapy/gocollyDemo/tvmao/web/"))
	http.Handle("/",  http.StripPrefix("",h))

	fmt.Println("Server is started")
}

func main() {

	startServer()
	//demo()

}

func startServer() {
	// 开启web服务器，并监听8092端口
	err := http.ListenAndServe("0.0.0.0:50001", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Start Server Running On IP localhost:50001")
}

func demo() {
	//搜索获取其相关信息
	key := "九州缥缈录"
	urlDrama := tvmao.Search(key)
	savePathDir := "E:/ZTestData/TVMao/"
	tvmao.StartCollect(urlDrama, savePathDir)
}
