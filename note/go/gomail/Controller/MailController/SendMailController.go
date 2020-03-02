/**
 * @Author: FB
 * @Description: 
 * @File:  SendMailController.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 17:06
 */

package MailController

import (
	"net/http"
	"strings"
	"gomail/Utils/Common"
	"fmt"
	"gomail/Server/SendMailServer"
	"gomail/Utils/Mail"
	"gomail/Utils/CommonVar"
)

func SendMailController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是二进制流

	switch r.Method {
	case "POST","GET":
		{
			//mailfrom := r.FormValue("mailfrom")
			mailto := r.FormValue("mailto")
			subject := r.FormValue("subject")
			content := r.FormValue("content")

			if strings.EqualFold(mailto, "") || strings.EqualFold(subject, "") || strings.EqualFold(content, "") {
				resp_json := CommonUtils.GetNormalRespJson(400, CommonVar.REQUEST_PARAMETER_INCORRECT, nil)
				fmt.Fprintln(w, string(resp_json))
				return
			}

			ok := MailUtils.CheckMailFormatValid(mailto)
			if !ok {
				resp_json := CommonUtils.GetNormalRespJson(400, CommonVar.INVALID_MAIL_FORMAT, nil)
				fmt.Fprintln(w, string(resp_json))
				return
			}

			ok,errcode := SendMailServer.SendMail(mailto,subject,content)
			if !ok {
				resp_json := CommonUtils.GetNormalRespJson(400, errcode, nil)
				fmt.Fprintln(w, string(resp_json))
				return
			}

			resp_json := CommonUtils.GetNormalRespJson(200, errcode, nil)
			fmt.Fprintln(w, string(resp_json))
			return
		}
	default:
		resp_json := CommonUtils.GetNormalRespJson(400, "错误的请求方式！", nil)
		fmt.Fprintln(w, string(resp_json))
		return
	}
}
