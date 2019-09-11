/**
 * @Author: FB
 * @Description: 模板信息等信息配置
 * @File:  SettingModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:08
 */

package Model

import "time"

type BasicSetting struct {
	ID        int    `gorm:"primary_key"`     // ID
	UUID      string `gorm:"not null;unique"` // XID
	Data      string                          // 数据信息
	Version   string                          // 信息版本信息
	Type      string                          // 类型,例如：模板信息，有效时间，消息推送模板
	CreatedAt time.Time                       // 创建时间
	UpdatedAt time.Time                       // 更新时间
	DeletedAt *time.Time `sql:"index"`        // 删除时间
}

func (BasicSetting) TableName() string {
	return "sd_basic_setting"
}
