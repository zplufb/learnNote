/**
 * @Author: FB
 * @Description: 
 * @File:  FormatBody.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 18:10
 */

package MailUtils

import (
	"strings"
	"fmt"
)

func FormatBody(body, verifyCode, expiredTime string) (bodyFormated string) {

	//您好！ 为确保是您本人操作，请在邮件验证码输入框输入下方验证码： {.verifyCode} 请勿向任何人泄露您收到的验证码。验证码有效期：{.expiredTime} 此致 鲜动定制
	newBody := strings.Replace(body, "{.verifyCode}", verifyCode, 1)
	newBody2 := strings.Replace(newBody, "{.expiredTime}", expiredTime+"分钟", 1)
	bodyFormated = strings.Replace(newBody2, "\b", "\n", -1)

	fmt.Println("bodyFormated=",bodyFormated)

	return
}
