/**
 * @Author: FB
 * @Description: 
 * @File:  GetMailSettingList.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 9:13
 */

package MailSettingServer

import (
	"gomail/Dao"
	"log"
	"gomail/Utils/CommonVar"
	"gomail/Model"
)

func GetMailSettingListInit(limit, offset int) (ok bool, errcode int, items []Model.MailSetting, total int) {

	var postCon Dao.Connect
	var postQuery Dao.Query
	var err error
	//var errcode int

	db, err := postCon.GetConnectionForPostgreSQL()
	if err != nil {
		errcode = CommonVar.DB_CONNECT_FAIL
		log.Println(err, errcode)
		return
	}

	// 获取邮箱账号设置列表
	items, err, total = postQuery.GetMailSettingList(db, limit, offset)
	if err != nil {
		errcode = CommonVar.MAIL_SETTING_LIST_FAIL
		log.Println(err, errcode)
		return
	}


	errcode = CommonVar.MAIL_SETTING_LIST_SUCCEE
	ok = true

	return

}
