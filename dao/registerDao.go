package dao

import (
	"fmt"
	"os/user"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

func InsertUser(user) {
	db, err := gorm.Open(
		mysql.Open("hitwh:hitwh1970@tcp(118.31.51.147:3306)/douyin")
	)
	//插入数据库
	
}