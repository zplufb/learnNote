/**
 * @Author: FB
 * @Description: 
 * @File:  CreateMailSettingController.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 9:41
 */

package BasicSettingController

import (
	"net/http"
	"strings"
	"gomail/Utils/Common"
	"fmt"
		"gomail/Server/BasicSettingServer"
	"gomail/Utils/CommonVar"
)

func CreateBasicSettingController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Origin", "*")             // 允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") // header的类型
	w.Header().Set("content-type", "application/json")             // 返回数据格式是二进制流

	switch r.Method {
	case "POST", "GET":
		{

			//mailfrom := r.FormValue("mailfrom")
			type_ := r.FormValue("type")
			data := r.FormValue("data")
			version := r.FormValue("version")

			if strings.EqualFold(type_, "") || strings.EqualFold(data, "") {
				resp_json := CommonUtils.GetNormalRespJson(400, CommonVar.REQUEST_PARAMETER_INCORRECT, nil)
				fmt.Fprintln(w, string(resp_json))
				return
			}

			if version == "" {
				version = "1.0"
			}

			ok, errcode := BasicSettingServer.CreateBasicSetting(data, type_,version)
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
