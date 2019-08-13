/**
 * @Author: FB
 * @Description: 下载连载小说
 * @File:  ReadNovel.go
 * @Version: 1.0.0
 * @Date: 2019/7/31 16:06
 */

package kanunu

import (
	"os"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"strconv"
	"math/rand"
	"time"
	"github.com/mahonia"
	"bufio"
)

//TODO 待修改
var Domain = "www.kanunu8.com"
//江南的九州缥缈录
var UrlIndex = "https://www.kanunu8.com/book2/11082/index.html"
var UrlInit = "https://www.kanunu8.com/book/4493/57833.html"
//var QueryTitle = "body > div > table:nth-child(4) > tbody > tr > td font"
var QueryTitle = "body > div > table:nth-child(4)  font"
//var QueryContent = "p:nth-child(1)"
var QueryContent = "body > div > table:nth-child(5)  p"
var page = 57833
var page_end = 57891

var Count int
var NovalText string
var fileName string
var fileName_href string

func init() {
	//删除指定文件
	fileName = `E:/novel_jn.txt`
	err := os.Remove(fileName)
	if err != nil {
		fmt.Errorf("err=%v", err)
	}

}



func CollectLinks(urlIndex string) (hrefMap map[string][]string) {

	hrefMap = make(map[string][]string)

	fileName_href = "E:/novel_chapter_links_tmp.txt"

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".book > dl", func(e *colly.HTMLElement) {


		//获取links
		contentArr := e.ChildAttrs("dd a", "href")
		len_hrefArr := len(contentArr)
		hrefArr := make([]string, len_hrefArr)

		for i, v := range contentArr {
			content := ConvertToString(v, "gbk", "utf-8")
			hrefArr[i] = content
		}
		hrefStr := strings.Join(hrefArr, ",")
		fmt.Println(hrefStr)

		//写到本地文本
		AppendToFile(fileName_href, hrefStr)
		hrefMap["a"] = hrefArr

		len_ := e.DOM.Find("dt").Length()
		//获取标题
		for i := 0; i < len_; i++ {
			str := e.DOM.Find("dt").Eq(i).Text()
			titleArr_ := ConvertToString(str, "gbk", "utf-8")
			fmt.Println("len_=", i, "->", titleArr_)

			strText := e.DOM.Find("dt").Eq(i).NextUntil("dt").Find("a[href]").Text()
			strText_ := ConvertToString(strText, "gbk", "utf-8")
			fmt.Println("len_=", i, "strText_->", strText_)

			arr := e.DOM.Find("dt").Eq(i).NextUntil("dt").Nodes
			for i, v := range arr {
				content := ConvertToString(v.Data, "gbk", "utf-8")
				fmt.Println("i=",i,"content=",content)
			}

		}



		//len_titleArr := len(titleArr)
		//titleSaveArr := make([]string, len_titleArr)
		//for i, v := range titleArr {
		//	content := ConvertToString(v, "gbk", "utf-8")
		//	titleSaveArr[i] = content
		//}
		//titleStr := strings.Join(titleSaveArr, ",")
		//fmt.Println(titleStr)
		//
		////写到本地文本
		//appendToFile(fileName_href, titleStr)
		//hrefMap["b"] = hrefArr

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) App leWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53"
	c.Visit(urlIndex)

	return
}

func ReadNovelDemo1() {

	//init
	page = 57867
	page_end = 57891
	UrlInit = "https://www.kanunu8.com/book/4493/" + strconv.Itoa(page) + ".html"
	pageNum := page_end - page
	fmt.Printf("Page From %v->%v pageNum=%v\n", page, page_end, pageNum)

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(Domain),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("body > div", func(e *colly.HTMLElement) {

		title_ := e.ChildText("table:nth-child(4)  font")
		content_ := e.ChildText("table:nth-child(5)  p")
		//fmt.Printf("Link found: %v\n", e.Text)
		title := ConvertToString(title_, "gbk", "utf-8")
		content := ConvertToString(content_, "gbk", "utf-8")
		//fmt.Printf("Link found: %v\n%v\n", title, content)

		//写到本地文本
		AppendToFile("E:/novel_jn.txt", "\n\n\n"+title+"\n\n")
		AppendToFile("E:/novel_jn.txt", content)

		Count++
		if Count >= pageNum {
			fmt.Println("End ")
		} else {
			randNum := rand.Int63n(1000)
			time.Sleep(time.Millisecond*300 + time.Duration(randNum))
			fmt.Println("Count=", Count, "time=", time.Now().Format("2006-01-02 15:04:05"))
			link := GetNextUrl()
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) App leWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53"
	c.Visit(UrlInit)
}

func GetNextUrl() (url_next string) {
	page ++
	str_page := strconv.Itoa(page)
	url_next = "https://www.kanunu8.com/book/4493/" + str_page + ".html"

	return
}

func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)

	return result

}

func AppendToFile(file, str string) (err error) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	num, err := writer.WriteString(str)
	fmt.Println("num=", num)
	err = writer.Flush()

	return

}
