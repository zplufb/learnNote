/**
 * @Author: FB
 * @Description: 
 * @File:  InitModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:05
 */

package InitServer

import (
	"gomail/Model"
	"gomail/Dao"
)

func InitModel() (err error) {

	// 从连接池获取数据库连接
	var postCon Dao.Connect
	db, err := postCon.GetConnectionForPostgreSQL()
	defer db.Close()

	// 基本配置信息
	if !db.HasTable("sd_basic_setting") {
		db.CreateTable(Model.BasicSetting{})
	}
	db.AutoMigrate(Model.BasicSetting{})

	// 邮箱账号配置信息
	if !db.HasTable("sd_mail_setting") {
		db.CreateTable(Model.MailSetting{})
	}
	db.AutoMigrate(Model.MailSetting{})

	// 验证码数据表
	if !db.HasTable("sd_verify_code") {
		db.CreateTable(Model.VerifyCode{})
	}
	db.AutoMigrate(Model.VerifyCode{})

	// 发邮件数据表
	if !db.HasTable("sd_send_mail") {
		db.CreateTable(Model.SendMail{})
	}
	db.AutoMigrate(Model.SendMail{})

	// 邮箱用户数据表
	if !db.HasTable("sd_user") {
		db.CreateTable(Model.User{})
	}
	db.AutoMigrate(Model.User{})

	// 组别数据表
	if !db.HasTable("sd_group") {
		db.CreateTable(Model.Group{})
	}
	db.AutoMigrate(Model.Group{})

	return

}
