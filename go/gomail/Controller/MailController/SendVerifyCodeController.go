/**
 * @Author: FB
 * @Description: 
 * @File:  SendVerifyCodeController.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:29
 */

package MailController

import (
	"net/http"
	"strings"
	"fmt"
	"gomail/Utils/Common"
	"gomail/Server/SendMailServer/VerifyCode"
	"gomail/Utils/CommonVar"
	"gomail/Utils/Mail"
)

func SendVerifyCodeController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是二进制流

	switch r.Method {
	case "POST","GET":
		{
			mailto := r.FormValue("mailto")
			type_ := r.FormValue("type")

			if strings.EqualFold(mailto, "") {
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

			ok,errcode := VefifyCodeServer.SendVerifyCode(mailto,type_)
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
