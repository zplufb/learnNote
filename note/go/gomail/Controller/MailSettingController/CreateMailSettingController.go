/**
 * @Author: FB
 * @Description: 
 * @File:  CreateMailSettingController.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 9:41
 */

package MailSettingController

import (
		"net/http"
	"strings"
	"gomail/Utils/Common"
	"fmt"
	"gomail/Server/MailSettingServer"
	"gomail/Utils/CommonVar"
)

func CreateMailSettingController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是二进制流

	switch r.Method {
	case "POST", "GET":
		{
			//User      string                          // 账号，例如：admin@shinedone.cn,Support@shinedone.cn
			//UserAlias string                          // 账号别称，例如：鲜动科技
			//Password  string                          // 密码
			//Host      string                          // 邮箱SMTP地址
			//Port      string                          // 端口
			//mailfrom := r.FormValue("mailfrom")
			user := r.FormValue("user")
			userAlias := r.FormValue("userAlias")
			password := r.FormValue("password")
			host := r.FormValue("host")
			port := r.FormValue("port")

			//port默认为465，可选是587
			if port == "" {
				port = "465"
			} else if port != "465" || port != "587" {
				port = "465"
			}

			if strings.EqualFold(user, "") || strings.EqualFold(password, "") || strings.EqualFold(host, "") {
				resp_json := CommonUtils.GetNormalRespJson(400, CommonVar.REQUEST_PARAMETER_INCORRECT, nil)
				fmt.Fprintln(w, string(resp_json))
				return
			}
			ok, errcode := MailSettingServer.CreateMailSetting(user, userAlias, password,host,port)
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
