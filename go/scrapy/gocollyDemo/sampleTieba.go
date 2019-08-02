/**
 * @Author: FB
 * @Description: 下载贴吧图片
 * @File:  sampleTieba.go
 * @Version: 1.0.0
 * @Date: 2019/8/1 16:54
 */

package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"log"
	"strconv"
)

func main() {

	url := "https://tieba.baidu.com/p/4803144798"
	pageNum := GetTiebaPage(url)
	needNum := 10
	GetTiebaImages(url, pageNum, needNum)
}

func GetTiebaImages(url string, pageNum int, needNum int) {

	if needNum > pageNum {
		fmt.Println("needNum 不能大于该帖总页数 pageNum", needNum, pageNum)
		return
	}

	//init
	//"https://tieba.baidu.com/p/6205499930" -> 6205499930
	subDirname := url[26:]
	saveFileName := "E:/tiebaImghref" + subDirname + ".txt"
	saveImgPath := "E:/images/tieba/" + subDirname + "/"
	var count = 1

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("tieba.baidu.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("#container", func(e *colly.HTMLElement) {

		//获取links
		contentArr := e.ChildAttrs(".left_section img.BDE_Image", "src")
		//len_hrefArr := len(contentArr)
		//hrefArr := make([]string, len_hrefArr)

		//贴吧不需要转换，因img中src没有带中文字符
		//for i, v := range contentArr {
		//	content := commonUtils.ConvertToString(v, "gbk", "utf-8")
		//	hrefArr[i] = content
		//}
		//hrefStr := strings.Join(hrefArr, ",")
		//fmt.Println(hrefStr)

		hrefStr := strings.Join(contentArr, ",")

		//写到本地文本
		AppendToFile(saveFileName, hrefStr)

		//下载到本地
		for _, v := range contentArr {
			err := DownloadImage(v, saveImgPath)
			if err != nil {
				log.Println(err)
				return
			}
		}

		count ++
		if needNum < count {
			fmt.Println("collect complete!")
			return
		}
		//获取下一个链接url
		nextUrl := getNextUrlForTieba(url, count)
		c.Visit(nextUrl)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) App leWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53"
	c.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	c.Visit(url)

}

func GetTiebaPage(url string) (pageNum int) {

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("tieba.baidu.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".pb_footer ul.l_posts_num", func(e *colly.HTMLElement) {

		//获取该帖的总长度
		sel := e.DOM.Find("li").First().Find("a").Last()
		content, _ := sel.Attr("href")

		pos_ := strings.LastIndex(content, "pn=")
		pageNumStr := content[pos_+3:]
		pageNum, _ = strconv.Atoi(pageNumStr)
		fmt.Println("content=", content, "pageNum=", pageNum)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("GetTiebaPage Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		//respContent := r.Body
		fmt.Println("OnResponse")
		//fmt.Println("respContent",string(respContent))
	})

	c.OnError(func(resp *colly.Response, err error) {
		fmt.Println("OnError")
	})

	c.OnScraped(func(resp *colly.Response) {
		//resp.Save("E:/123.txt")
		fmt.Println("ScrapedCallback")
	})

	//c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) App leWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53"
	c.Visit(url)

	return

}

func getNextUrlForTieba(url string, pageNo int) (nextUrl string) {

	pageNoStr := strconv.Itoa(pageNo)
	nextUrl = url + "?pn=" + pageNoStr

	return
}
