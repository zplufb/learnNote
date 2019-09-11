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
)

func GetMailSettingList(limit, offset int) (ok bool, errcode int, resp_map []map[string]interface{}, total int) {

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
	data, err, total := postQuery.GetMailSettingList(db, limit, offset)
	if err != nil {
		errcode = CommonVar.MAIL_SETTING_LIST_FAIL
		log.Println(err, errcode)
		return
	}

	loop := len(data)
	resp_map = make([]map[string]interface{}, loop, loop)
	for i := 0; i < loop; i++ {
		item := make(map[string]interface{})
		item["id"] = data[i].UUID
		item["user"] = data[i].User
		item["password"] = data[i].Password
		item["host"] = data[i].Host
		item["port"] = data[i].Port
		resp_map[i] = item
	}

	errcode = CommonVar.MAIL_SETTING_LIST_SUCCEE
	ok = true

	return

}
