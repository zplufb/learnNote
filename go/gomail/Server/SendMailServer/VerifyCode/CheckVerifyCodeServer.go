/**
 * @Author: FB
 * @Description: 
 * @File:  GetVerifyCodeServer.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 11:26
 */

package VefifyCodeServer

import (
	"gomail/Utils/CommonVar"
	"log"
	"gomail/Dao"
	"time"
)

func CheckVerifyCode(mailto string,verifyCode string) (ok bool, errcode int) {

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
	record, err := postQuery.GetVerifyCodeRecord(db, mailto)
	if err != nil {
		errcode = CommonVar.VERIFY_CODE_GET_RECORD_FAIL
		log.Println(err, errcode)
		return
	}

	verifyCodeServer := record.VerifyCode
	expiredTime := record.ExpiredTime

	now := time.Now()
	flag_no_expired := now.Before(expiredTime)

	//验证码还未过期
	if flag_no_expired {
		if verifyCodeServer != verifyCode {
			errcode = CommonVar.VERIFY_CODE_CHECK_FAIL
			return
		}
	}else{
		errcode = CommonVar.VERIFY_CODE_TIMEOUT
		return
	}


	errcode = CommonVar.VERIFY_CODE_CHECK_SUCCEE
	ok = true

	return

}
