/**
 * @Author: FB
 * @Description: 用户验证验证码
 * @File:  VerifyCodeModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:08
 */

package Model

import "time"

type VerifyCode struct {
	ID          int    `gorm:"primary_key"`     // ID
	UUID        string `gorm:"not null;unique"` // XID
	MailFrom    string                          // 邮件发送人，默认为官方帐号
	MailTo      string                          // 收件人邮箱地址
	VerifyCode  string                          // 验证码
	Type        string                          // 类型，例如：首次绑定，修改密码，注销账号
	ExpiredTime time.Time                       // 截止日期
	CreatedAt   time.Time                       // 创建时间
	UpdatedAt   time.Time                       // 更新时间
	DeletedAt   *time.Time `sql:"index"`        // 删除时间
}

func (VerifyCode) TableName() string {
	return "sd_verify_code"
}
