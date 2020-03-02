/**
 * @Author: FB
 * @Description: 
 * @File:  MailUtils
 * @Version: 1.0.0
 * @Date: 2019/9/9 14:31
 */

package MailUtils

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"regexp"
)

func SendMail(mailConn map[string]string, mailTo []string, subject string, body string) (err error) {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	//mailConn := map[string]string{
	//	"user": "zplufb@163.com",
	//	"pass": "lfb@0729.com",
	//	"host": "smtp.163.com",
	//	"port": "465",
	//}

	//TODO test
	return nil

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], mailConn["alias"])) //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", FromNameAlias+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err = d.DialAndSend(m)
	return err

}

func CheckMailFormatValid(emails ...string) bool {

	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)

	var flag = false
	for _, email := range emails {
		flag = reg.MatchString(email)
		if !flag {
			return false
		}
	}

	return true

}
