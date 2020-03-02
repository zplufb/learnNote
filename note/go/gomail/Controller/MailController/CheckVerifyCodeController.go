/**
 * @Author: FB
 * @Description: 
 * @File:  GetVerifyCodeController.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 11:24
 */

package MailController

import (
	"net/http"
	"strings"
	"gomail/Utils/Common"
	"fmt"
	"gomail/Server/SendMailServer/VerifyCode"
	"gomail/Utils/CommonVar"
	"gomail/Utils/Mail"
)

func CheckVerifyCodeController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是二进制流

	switch r.Method {
	case "POST", "GET":
		{
			mailto := r.FormValue("mailto")
			verifycode := r.FormValue("verifycode")

			if strings.EqualFold(mailto, "") || strings.EqualFold(verifycode, "") {
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

			ok, errcode := VefifyCodeServer.CheckVerifyCode(mailto, verifycode)
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
