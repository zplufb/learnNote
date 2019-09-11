/**
 * @Author: FB
 * @Description: 
 * @File:  ServerInit.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:27
 */

package ServerInit

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "gomail/Server/Init"
	"net/http"
	"os"
	"strings"
	"fmt"
	"gomail/Controller/MailController"
	"gomail/Utils/Common"
	"gomail/Controller/MailSettingController"
	"gomail/Controller/BasicSettingController"
	"gomail/Server/Init"
)

func init() {
	fmt.Println("Server is starting...")

	// 初始化所有对象和变量
	InitServer.InitModel()
	InitServer.InitVar()

	// 启动请求路由
	route()
	// 启动静态文件服务
	fileServer()
	fmt.Println("Server has started...")
}

func route() {
	// 发送验证码
	http.HandleFunc("/sendVerifyCode", MailController.SendVerifyCodeController)
	http.HandleFunc("/checkVerifyCode", MailController.CheckVerifyCodeController)

	// 发送邮件
	http.HandleFunc("/sendMail", MailController.SendMailController)

	//配置邮箱账号信息
	http.HandleFunc("/createMailSetting", MailSettingController.CreateMailSettingController)
	http.HandleFunc("/getMailSettingList", MailSettingController.GetMailSettingList)
	//
	////基本配置信息
	http.HandleFunc("/getBasicSettingList", BasicSettingController.GetBasicSettingList)
	http.HandleFunc("/createBasicSetting", BasicSettingController.CreateBasicSettingController)

}

func fileServer() {
	if CommonUtils.IsWindowsPlatform() {
		h := http.FileServer(http.Dir("/Go-Work/src/gomail/web/"))
		http.Handle("/", http.StripPrefix("", protect(h)))
	} else {
		host, _ := os.Hostname()
		if strings.EqualFold(host, "ushirts1G6") {
			//ubuntu
			h := http.FileServer(http.Dir("/usr/local/Go-Work/src/gomail/web/"))
			http.Handle("/", http.StripPrefix("", protect(h)))
		}
	}
}

// 目录访问保护
func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		end := url[len(url)-1 : len(url)]
		if len(url) > 4 {
			js := url[len(url)-2 : len(url)]
			css := url[len(url)-3 : len(url)]
			// 如果请求访问的js和css文件，则设置缓存一星期
			if strings.EqualFold(js, "js") {
				w.Header().Set("Cache-Control", "max-age=86400")
			}
			if strings.EqualFold(css, "css") {
				w.Header().Set("Cache-Control", "max-age=86400")
			}
			if strings.EqualFold(css, "png") {
				w.Header().Set("Cache-Control", "max-age=86400")
				//w.Header().Set("Cache-Control", "no-cache")
			}
		}

		// 如果请求访问的是目录列表，直接跳转到首页
		if strings.EqualFold(end, "/") && len(url) > 1 {
			redirect := "https://" + r.Host + "/"
			http.Redirect(w, r, redirect, 302)
		}

		h.ServeHTTP(w, r)
	})
}
