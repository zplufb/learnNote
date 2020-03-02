/**
 * @Author: FB
 * @Description: 
 * @File:  GroupModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 10:51
 */

package Model

import "time"

type Group struct {
	ID        int    `gorm:"primary_key"`     // ID
	UUID      string `gorm:"not null;unique"` // XID
	Name      string                          // 组别名称
	Etc       string                          // 备注
	CreatedAt time.Time                       // 创建时间
	UpdatedAt time.Time                       // 更新时间
	DeletedAt *time.Time `sql:"index"`        // 删除时间
}

func (Group) TableName() string {
	return "sd_group"
}
