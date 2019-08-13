package tvmao

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
	"time"
	"Test001/golang/scrapy/gocollyDemo/utils"
	"strconv"
	"net/url"
	"log"
)

//自动在该网站搜索电视剧名称
var http_proxy = "https://"
var domain = "www.tvmao.com"

func Search(key string) (href string) {
	//时间都知道=>%E6%97%B6%E9%97%B4%E9%83%BD%E7%9F%A5%E9%81%93
	keys := url.PathEscape(key)
	urlInit := "https://www.tvmao.com/query.jsp?keys=" + keys

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("#t_q_tab_drama", func(e *colly.HTMLElement) {
		var exist bool
		//TODO 默认取第一个链接
		href, exist = e.DOM.Find(".rtive").First().Find("a").Attr("href")
		if exist {
			fmt.Println("href=", href)
		}

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Search Visiting", r.URL.String())
	})

	//UserAgent PC
	c.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"
	c.Visit(urlInit)

	return

}

func GetInfo(urlInit string) (title string, info string, epiCount int) {

	// 获取剧情内容
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML(".section-wrap", func(e *colly.HTMLElement) {

		title = e.DOM.Find("strong.font20").Text()
		info = e.DOM.Find(".obj_meta").Text()

		num := e.DOM.Find(".obj_meta tr >td").Eq(1).Find("span").Text()
		epiCount, _ = strconv.Atoi(num)
		//fmt.Println("title,info,pageNum=",title,info,pageNum)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("GetInfo Visiting", r.URL.String())
	})

	//UserAgent IPhone
	//c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) App leWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53"
	//UserAgent PC
	c.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"
	c.Visit(urlInit)

	return
}
func StartCollect(urlDrama string,savePathDir string) {

	//获取标题，信息，集数
	urlInit := http_proxy + domain + urlDrama
	title, info, epiCount := GetInfo(urlInit)
	//判定文件夹是否存在，若不存在则新建
	ok ,err := CollyUtils.AutoCreateDirWhenNotExist(savePathDir)
	if !ok{
		log.Println(err)
		return
	}


	//写到本地文本
	fileName := savePathDir + title + ".txt"
	CollyUtils.AppendToFile(fileName, info)
	//init
	urlEpisode := http_proxy + domain + urlDrama + "/episode"
	fmt.Println("title=", title, "\nInfo=", info, "\nepiCount=", epiCount, "\nfilename=", fileName, "\ninitUrl=", urlInit, "\nurlEpisode=", urlEpisode)

	//init var
	count := 0

	//选择器
	//$("strong.font20").text() //标题
	//$(".obj_meta").text() 获取电视剧基本信息
	//$(".obj_meta tr:first > td:nth-child(2) >span").text() //获取电视剧的总集数

	//$("dl.menu_tab.clear > dd:nth-child(2) > a").attr("href") //获取分集剧情的href
	//$(".epi_t").text() //获取单集的标题
	//$("article.epi_c p").text() //获取单集剧情
	//$("article +.alignct a").attr("href") //获取第二集的href
	//$("article +.alignct a:nth-child(2)").attr("href") ////获取下一集的href

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	// 获取剧情内容
	c.OnHTML(".section-wrap", func(e *colly.HTMLElement) {

		chapter_name := e.DOM.Find(".epi_t").Text()
		content := e.DOM.Find("article.epi_c p").Text()

		var href string
		href, _ = e.DOM.Find("article +.alignct a:nth-child(2)").Attr("href")
		if href == "" {
			//只有第一集的链接只有一个：下一集。
			href, _ = e.DOM.Find("article +.alignct a").Attr("href")
		}

		fmt.Printf("Next Chapter Link found: %v\n", href)

		//写到本地文本
		CollyUtils.AppendToFile(fileName, "\n\n"+chapter_name+"\n\n")
		CollyUtils.AppendToFile(fileName, content)

		count++
		if count >= epiCount {
			fmt.Println("Collect Completed ")
		} else {
			randNum := rand.Int63n(1000)
			time.Sleep(time.Millisecond*300 + time.Duration(randNum))
			fmt.Println("count=", count, "time=", time.Now().Format("2006-01-02 15:04:05"))
			link := href
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//UserAgent IPhone
	//c.UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) App leWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53"
	//UserAgent PC
	c.UserAgent = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36"
	c.Visit(urlEpisode)

}
