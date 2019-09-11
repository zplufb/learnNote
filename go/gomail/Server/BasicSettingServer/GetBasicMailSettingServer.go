/**
 * @Author: FB
 * @Description: 
 * @File:  GetBasicMailSettingServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 17:24
 */

package BasicSettingServer

import (
	"gomail/Utils/CommonVar"
	"log"
	"gomail/Dao"
	"gomail/Model"
)

func GetBasicSettingById(id int) (ok bool, errcode int, data Model.BasicSetting) {

	var postCon Dao.Connect
	var postQuery Dao.Query
	var err error

	db, err := postCon.GetConnectionForPostgreSQL()
	if err != nil {
		errcode = CommonVar.DB_CONNECT_FAIL
		log.Println(err, errcode)
		return
	}

	// 获取基本配置记录
	data, err = postQuery.GetBasicSettingById(db, id)
	if err != nil {
		errcode = CommonVar.BASIC_SETTING_GET_RECORD_FAIL
		log.Println(err, errcode)
		return
	}

	errcode = CommonVar.BASIC_SETTING_GET_RECORD_SUCCEE
	ok = true

	return

}
