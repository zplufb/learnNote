/**
 * @Author: FB
 * @Description: 
 * @File:  UserModel.go
 * @Version: 1.0.0
 * @Date: 2019/9/10 10:51
 */

package Model

import "time"

type User struct {
	ID        int    `gorm:"primary_key"`     // ID
	UUID      string `gorm:"not null;unique"` // XID
	Name      string                          // 用户
	Group     string                          // 组别，默认是用户组，还有VIP，设计师等
	CreatedAt time.Time                       // 创建时间
	UpdatedAt time.Time                       // 更新时间
	DeletedAt *time.Time `sql:"index"`        // 删除时间
}

func (User) TableName() string {
	return "sd_user"
}
