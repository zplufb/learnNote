/**
 * @Author: FB
 * @Description: 
 * @File:  Query.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:02
 */

package Dao

import (
	"github.com/jinzhu/gorm"
	"gomail/Model"
)

type Query struct {
}

////SendMail
///验证码
func (*Query) GetVerifyCodeRecord(db *gorm.DB, mailto string) (item Model.VerifyCode, err error) {
	err = db.Table("sd_verify_code").Where("mail_to=?", mailto).Last(&item).Error
	return
}

///发邮件列表
func (*Query) GetMailRecordList(db *gorm.DB, limit int, offset int) (items []Model.SendMail, err error, total int) {
	err = db.Table("sd_send_mail").Order("id").Limit(limit).Offset(offset).Find(&items).Error
	db.Raw("select count(*) - (select count(*) from sd_mail_setting where deleted_at > 'epoch') from sd_send_mail").Row().Scan(&total)
	return
}

////MailSetting
func (*Query) GetMailSettingList(db *gorm.DB, limit int, offset int) (items []Model.MailSetting, err error, total int) {
	err = db.Table("sd_mail_setting").Order("id").Limit(limit).Offset(offset).Find(&items).Error
	db.Raw("select count(*) - (select count(*) from sd_mail_setting where deleted_at > 'epoch') from sd_mail_setting").Row().Scan(&total)
	return
}

////BasicSetting
func (*Query) GetBasicSettingList(db *gorm.DB, limit int, offset int) (items []Model.BasicSetting, err error, total int) {
	err = db.Table("sd_basic_setting").Order("id").Limit(limit).Offset(offset).Find(&items).Error
	db.Raw("select count(*) - (select count(*) from sd_basic_setting where deleted_at > 'epoch') from sd_basic_setting").Row().Scan(&total)
	return
}

func (*Query) GetBasicSettingById(db *gorm.DB, id int) (item Model.BasicSetting, err error) {
	err = db.Table("sd_basic_setting").Where("id=?",id).First(&item).Error
	return
}

////User
func (*Query) GetUserList(db *gorm.DB, limit int, offset int) (items []Model.User, err error, total int) {
	err = db.Table("sd_user").Order("id").Limit(limit).Offset(offset).Find(&items).Error
	db.Raw("select count(*) - (select count(*) from sd_user where deleted_at > 'epoch') from sd_user").Row().Scan(&total)
	return
}

////组别
func (*Query) GetGroupList(db *gorm.DB, limit int, offset int) (items []Model.Group, err error, total int) {
	err = db.Table("sd_group").Order("id").Limit(limit).Offset(offset).Find(&items).Error
	db.Raw("select count(*) - (select count(*) from sd_group where deleted_at > 'epoch') from sd_group").Row().Scan(&total)
	return
}
