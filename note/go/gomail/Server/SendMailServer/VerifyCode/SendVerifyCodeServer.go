/**
 * @Author: FB
 * @Description: 
 * @File:  SendVerifyCodeServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:34
 */

package VefifyCodeServer

import (
	"gomail/Utils/Mail"
	"fmt"
	"log"
	"gomail/Utils/CommonVar"
	"gomail/Utils/UUID"
)

func SendVerifyCode(mailAddress, type_ string) (ok bool, errcode int) {

	//定义收件人
	mailTo := []string{
		mailAddress,
	}

	b, _ := UUID.Base64Decode(CommonVar.MailSettingMap[CommonVar.TYPE_MS_ADMIN].Password)
	password := string(b)
	mailConn := map[string]string{
		"user": CommonVar.MailSettingMap[CommonVar.TYPE_MS_ADMIN].User,
		"alias":CommonVar.MailSettingMap[CommonVar.TYPE_MS_ADMIN].UserAlias,
		"pass": password,
		"host": CommonVar.MailSettingMap[CommonVar.TYPE_MS_ADMIN].Host,
		"port": CommonVar.MailSettingMap[CommonVar.TYPE_MS_ADMIN].Port,
	}

	//邮件主题为"Hello"
	//subject := "[VerifyCode]Ushirt邮箱绑定验证码"
	subject := CommonVar.BasicSettingMap[CommonVar.TYPE_BS_SUBJECT].Data
	// 邮件正文
	verifyCode := UUID.GetRandCode(6)
	//body := "VerifyCode: " + verifyCode + ". It will expire after 5 minute,Thanks!"
	body0 := CommonVar.BasicSettingMap[CommonVar.TYPE_BS_BODY].Data
	expired := CommonVar.BasicSettingMap[CommonVar.TYPE_BS_TIME].Data

	body := MailUtils.FormatBody(body0, verifyCode, expired)

	err := MailUtils.SendMail(mailConn, mailTo, subject, body)
	if err != nil {
		errcode = CommonVar.VERIFY_CODE_SEND_FAIL
		log.Println("send fail,err=", err)
		return
	}

	//插入验证码数据库
	mailFrom := mailConn["user"]

	ok, errcode = CreateVerifyCode(mailFrom, mailAddress, verifyCode, type_)
	if !ok {
		return
	}

	ok = true
	errcode = CommonVar.VERIFY_CODE_SEND_SUCCEE

	fmt.Println("send to " + mailAddress + " successfully")

	return
}
