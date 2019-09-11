/**
 * @Author: FB
 * @Description: 用户验证验证码
 * @File:  VerifyCodeModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:08
 */

package Model

import "time"

type SendMail struct {
	ID        int    `gorm:"primary_key"`     // ID
	UUID      string `gorm:"not null;unique"` // XID
	MailFrom  string                          // 邮件发送人
	MailTo    []User                          // 收件人邮箱地址数组
	Subject   string                          // 标题
	Body      string                          // 正文内容
	Type      string                          // 类型，例如：消息推送pushMsg（VIP，设计师，所有用户等群组）
	CreatedAt time.Time                       // 创建时间
	UpdatedAt time.Time                       // 更新时间
	DeletedAt *time.Time `sql:"index"`        // 删除时间
}

func (SendMail) TableName() string {
	return "sd_send_mail"
}
