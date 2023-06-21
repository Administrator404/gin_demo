package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// 连接数据库
func OpenDatabase() {
	dsn := "root:J21kw083z547@tcp(hyh-yys.top:3306)/ginDemo01?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接异常：", err)
	}
}
