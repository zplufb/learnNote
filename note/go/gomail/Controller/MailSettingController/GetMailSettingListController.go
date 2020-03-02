/**
 * @Author: FB
 * @Description: 
 * @File:  MailSettingList
 * @Version: 1.0.0
 * @Date: 2019/9/10 9:09
 */

package MailSettingController

import (
	"net/http"
	"gomail/Utils/Common"
	"fmt"
	"strconv"
	"gomail/Server/MailSettingServer"
)

func GetMailSettingList(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是二进制流

	switch r.Method {
	case "POST", "GET":
		{
			// 一页显示多少条数据
			limit, _ := strconv.Atoi(r.FormValue("limit"))
			// 当前页数
			page, _ := strconv.Atoi(r.FormValue("page"))

			//假定数据
			limit = 10
			page = 0

			ok, errcode, resp, total := MailSettingServer.GetMailSettingList(limit, page)
			if !ok {
				resp_json := CommonUtils.GetListRespJson(400, errcode, nil, 0)
				fmt.Fprintln(w, string(resp_json))
				return
			}

			resp_json := CommonUtils.GetListRespJson(200, errcode, resp, total)
			fmt.Fprintln(w, string(resp_json))
			return
		}
	default:
		resp_json := CommonUtils.GetListRespJson(400, "错误的请求方式！", nil, 0)
		fmt.Fprintln(w, string(resp_json))
		return
	}
}
