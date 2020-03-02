/**
 * @Author: FB
 * @Description: 
 * @File:  Connect.go
 * @Version: 1.0.0
 * @Date: 2019/9/9 15:02
 */

package Dao

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Connect struct {
}

var (
	db  *gorm.DB
	err error
)

func (Connect) GetConnectionForPostgreSQL() (db *gorm.DB, err error) {
	if db != nil {
		return
	} else {
		db, err = gorm.Open("postgres",
			"host=192.168.3.168 port=5432 user=postgres dbname=mail password=123456 sslmode=disable")

		if err != nil {
			log.Println(err)
			panic("failed to connect database")
		}

		db.DB().SetMaxIdleConns(10) // 最大空闲连接数量 	10个
		db.DB().SetMaxOpenConns(20) // 最大连接数量		20个
	}

	return
}
