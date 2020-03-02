/**
 * @Author: FB
 * @Description: 
 * @File:  SendMailServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 17:08
 */

package SendMailServer

import (
	"gomail/Utils/Mail"
	"gomail/Utils/CommonVar"
	"fmt"
)

func SendMail(mailto, subject, body string) (ok bool, errcode int) {

	//定义收件人
	mailTo := []string{
		mailto,
	}

	//mailConn := map[string]string{
	//	"user": "lufb@shinedone.cn",
	//	"pass": "r4r3StCjCjH7a7Uk",
	//	"host": "smtp.exmail.qq.com",
	//	"port": "465",
	//}

	mailConn := map[string]string{
		"user": "lufb@shinedone.cn",
		"pass": "r4r3StCjCjH7a7Uk",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}

	err := MailUtils.SendMail(mailConn, mailTo, subject, body)
	if err != nil {
		errcode = CommonVar.SEND_MAIL_SEND_FAIL
		return
	}

	ok = true
	errcode = CommonVar.SEND_MAIL_SEND_SUCCEE

	fmt.Println("send to " + mailto + " successfully")
	return
}
