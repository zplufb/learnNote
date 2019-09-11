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
)

func GetBasicSettingList(limit, offset int) (ok bool, errcode int, resp_map []map[string]interface{}, total int) {

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
	data, err, total := postQuery.GetBasicSettingList(db, limit, offset)
	if err != nil {
		errcode = CommonVar.BASIC_SETTING_LIST_FAIL
		log.Println(err, errcode)
		return
	}

	loop := len(data)
	resp_map = make([]map[string]interface{}, loop, loop)
	for i := 0; i < loop; i++ {
		item := make(map[string]interface{})
		item["id"] = data[i].UUID
		item["data"] = data[i].Data
		item["type"] = data[i].Type
		item["version"] = data[i].Version
		resp_map[i] = item
	}

	errcode = CommonVar.BASIC_SETTING_LIST_SUCCEE
	ok = true

	return

}
