/**
 * @Author: FB
 * @Description: 邮件服务器信息配置
 * @File:  SettingModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:08
 */

package Model

import "time"

type MailSetting struct {
	ID        int    `gorm:"primary_key"`     // ID
	UUID      string `gorm:"not null;unique"` // XID
	User      string                          // 账号，例如：admin@shinedone.cn,Support@shinedone.cn
	UserAlias string                          // 账号别称，例如：鲜动科技
	Password  string                          // 密码
	Host      string                          // 邮箱SMTP地址
	Port      string                          // 端口
	CreatedAt time.Time                       // 创建时间
	UpdatedAt time.Time                       // 更新时间
	DeletedAt *time.Time `sql:"index"`        // 删除时间
}

func (MailSetting) TableName() string {
	return "sd_mail_setting"
}
