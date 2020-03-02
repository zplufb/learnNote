/**
 * @Author: FB
 * @Description: 
 * @File:  Insert.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:02
 */

package Dao

import (
	"github.com/jinzhu/gorm"
	"gomail/Model"
)

type Insert struct {
}

////SendMail
///验证码
func (*Insert) CreateVerifyCodeRecord(db *gorm.DB, item Model.VerifyCode) (err error) {
	err = db.Create(&item).Error
	return
}
///发邮件（单独或者群发）
func (*Insert) CreateMailRecord(db *gorm.DB, item Model.SendMail) (err error) {
	err = db.Create(&item).Error
	return
}

////MailSetting
func (*Insert) CreateMailSettingRecord(db *gorm.DB, item Model.MailSetting) (err error) {
	err = db.Create(&item).Error
	return
}

////BasicSetting
func (*Insert) CreateBasicSettingRecord(db *gorm.DB, item Model.BasicSetting) (err error) {
	err = db.Create(&item).Error
	return
}

////User，首次绑定时插入
func (*Insert) CreateUserRecord(db *gorm.DB, item Model.User) (err error) {
	err = db.Create(&item).Error
	return
}

////组别
func (*Insert) CreateGroupRecord(db *gorm.DB, item Model.Group) (err error) {
	err = db.Create(&item).Error
	return
}
