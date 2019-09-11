/**
 * @Author: FB
 * @Description: 
 * @File:  CreateVerifyCodeServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 12:51
 */

package VefifyCodeServer

import (
	"gomail/Utils/CommonVar"
	"log"
	"gomail/Model"
	"gomail/Dao"
	"gomail/Utils/UUID"
	"time"
	"strconv"
)

func CreateVerifyCode(mailFrom, mailTo, code, type_ string) (ok bool, errcode int) {
	var postCon Dao.Connect
	var postInsert Dao.Insert

	// 从连接池获取数据库连接
	db, err := postCon.GetConnectionForPostgreSQL()
	if err != nil {
		errcode = CommonVar.DB_CONNECT_FAIL
		log.Println(err)
		return
	}

	var item Model.VerifyCode

	item.UUID = UUID.GetXID()
	item.MailFrom = mailFrom
	item.MailTo = mailTo
	item.VerifyCode = code

	typeInt, _ := strconv.Atoi(type_)
	item.Type = CommonVar.TypeMap[typeInt]
	//5分钟
	d := time.Minute * 5
	item.ExpiredTime = time.Now().Add(d)

	err = postInsert.CreateVerifyCodeRecord(db, item)
	if err != nil {
		errcode = CommonVar.VERIFY_CODE_CREATE_FAIL
		log.Println(err)
		return
	}

	ok = true
	errcode = CommonVar.VERIFY_CODE_CREATE_SUCCEE

	return
}
