package main

import (
	"ajun/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("错误:", err.Error())
	}

	a := make([]model.Information, 0)
	db.Raw("INSERT INTO information (name,sex,age) VALUES('肖','男',30)").Scan(&a)
	db.Raw("SELECT * FROM information").Scan(&a)
	fmt.Println(a)

}
