/**
 * @Author: FB
 * @Description: 
 * @File:  CreateMailSettingServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 9:48
 */

package MailSettingServer

import (
	"log"
	"gomail/Model"
	"gomail/Utils/UUID"
	"gomail/Dao"
	"gomail/Utils/CommonVar"
)

func CreateMailSetting(user, userAlias, password, host, port string) (ok bool, errcode int) {
	var postCon Dao.Connect
	var postInsert Dao.Insert

	// 从连接池获取数据库连接
	db, err := postCon.GetConnectionForPostgreSQL()
	if err != nil {
		errcode = CommonVar.DB_CONNECT_FAIL
		log.Println(err)
		return
	}

	var item Model.MailSetting

	item.UUID = UUID.GetXID()
	item.User = user
	item.UserAlias = userAlias
	item.Password = UUID.Base64Encode(password)
	item.Host = host
	item.Port = port

	err = postInsert.CreateMailSettingRecord(db, item)
	if err != nil {
		errcode = CommonVar.MAIL_SETTING_CREATE_FAIL
		log.Println(err)
		return
	}

	ok = true
	errcode = CommonVar.MAIL_SETTING_CREATE_SUCCEE

	return
}
