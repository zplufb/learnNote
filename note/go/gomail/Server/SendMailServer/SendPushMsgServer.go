/**
 * @Author: FB
 * @Description: 
 * @File:  SendPushMsgServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 16:45
 */

package SendMailServer

import (
	"gomail/Utils/Mail"
			"fmt"
	"gomail/Utils/CommonVar"
)

func SendPushMsgServer(mailAddresses []string) (ok bool, errcode int) {

	mailConn := map[string]string{
		"user": "lufb@shinedone.cn",
		"pass": "r4r3StCjCjH7a7Uk",
		"host": "smtp.exmail.qq.com",
		"port": "465",
	}
	//定义收件人
	mailTo := mailAddresses
	//邮件主题为"Hello"
	subject := "Hello by golang gomail from exmail.qq.com"
	// 邮件正文
	body := "Hello,by gomail sent 33 all"

	err := MailUtils.SendMail(mailConn,mailTo, subject, body)
	if err != nil {
		errcode = CommonVar.SEND_MAIL_SEND_FAIL
		return
	}

	errcode = CommonVar.SEND_MAIL_SEND_SUCCEE

	for i, v := range mailTo {
		fmt.Println("send to "+v+" successfully,no=", i)
	}

	return
}
