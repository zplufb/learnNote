package searchController

import (
		"fmt"
	"net/http"
		"Test001/golang/scrapy/gocollyDemo/utils"
<<<<<<< HEAD
	"Test001/golang/scrapy/gocollyDemo/tvmao/server"
=======
>>>>>>> fd158c889244924c570e72ea81383ff0b0dfacd6
)

func Search(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// 允许访问所有域
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是JSON

	switch r.Method {
<<<<<<< HEAD
	case "GET","POST":
=======
	case "POST":
>>>>>>> fd158c889244924c570e72ea81383ff0b0dfacd6
		{
			key := r.FormValue("key")

			if key == ""{
				resp_json := CollyUtils.GetRespJson(400, "参数有误！", key)
				fmt.Fprintln(w, string(resp_json))
				return
			}

<<<<<<< HEAD
			ok := searchServer.StartSearch(key)

=======
			//ok := searchServer.StartSearch(key)
			ok := true
>>>>>>> fd158c889244924c570e72ea81383ff0b0dfacd6
			if ok {
				resp_json := CollyUtils.GetRespJson(200, "下载成功！", key)
				fmt.Fprintln(w, string(resp_json))
				return
			}
			resp_json := CollyUtils.GetRespJson(500, "服务器错误！", nil)
			fmt.Fprintln(w, string(resp_json))
			return
		}
	default:
		resp_json := CollyUtils.GetRespJson(300, "错误的请求方式！", nil)
		fmt.Fprintln(w, string(resp_json))
		return
	}
}
