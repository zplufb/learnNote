/**
 * @Author: FB
 * @Description: 
 * @File:  GetBasicSettingList.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 14:14
 */

package BasicSettingServer

import (
	"gomail/Utils/CommonVar"
	"log"
	"gomail/Dao"
	"gomail/Model"
)

func GetBasicSettingListInit(limit, offset int) (ok bool, errcode int, items []Model.BasicSetting, total int) {

	var postCon Dao.Connect
	var postQuery Dao.Query
	var err error

	db, err := postCon.GetConnectionForPostgreSQL()
	if err != nil {
		errcode = CommonVar.DB_CONNECT_FAIL
		log.Println(err, errcode)
		return
	}

	// 获取邮箱账号设置列表
	items, err, total = postQuery.GetBasicSettingList(db, limit, offset)
	if err != nil {
		errcode = CommonVar.BASIC_SETTING_LIST_FAIL
		log.Println(err, errcode)
		return
	}

	errcode = CommonVar.BASIC_SETTING_LIST_SUCCEE
	ok = true

	return

}
