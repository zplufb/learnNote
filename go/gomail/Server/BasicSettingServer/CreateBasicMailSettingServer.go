/**
 * @Author: FB
 * @Description: 
 * @File:  CreateMailSettingServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 9:48
 */

package BasicSettingServer

import (
	"log"
	"gomail/Model"
	"gomail/Utils/UUID"
	"gomail/Dao"
	"gomail/Utils/CommonVar"
	"strconv"
)

func CreateBasicSetting(data, type_,version string) (ok bool, errcode int) {
	var postCon Dao.Connect
	var postInsert Dao.Insert

	// 从连接池获取数据库连接
	db, err := postCon.GetConnectionForPostgreSQL()
	if err != nil {
		errcode = CommonVar.DB_CONNECT_FAIL
		log.Println(err)
		return
	}

	var item Model.BasicSetting

	//type映射
	typeInt,_ := strconv.Atoi(type_)
	typeStr := CommonVar.TypeMap[typeInt]


	item.UUID = UUID.GetXID()
	item.Data = data
	item.Type = typeStr
	item.Version = version

	err = postInsert.CreateBasicSettingRecord(db, item)
	if err != nil {
		errcode = CommonVar.BASIC_SETTING_CREATE_FAIL
		log.Println(err)
		return
	}

	ok = true
	errcode = CommonVar.BASIC_SETTING_CREATE_SUCCEE

	return
}
