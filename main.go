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
	b := string("information")
	c := "李狗蛋"
	d := "男"
	e := "35"
	a := make([]model.Information, 0)
	db.Raw("TRUNCATE TABLE " + b).Scan(&a)                                                               //清空表格所有信息
	db.Raw("SELECT *FROM information").Scan(&a)                                                          //查询表格所有信息
	db.Raw("INSERT INTO information(name,sex,age) VALUES  ('" + c + "','" + d + "'," + e + ")").Scan(&a) //增
	db.Raw("DELETE FROM information WHERE name='李'")                                                     //删
	db.Raw("UPDATE information SET ID='1',`name`='杨',sex='女',age=25 WHERE  `name`='肖'").Scan(&a)         //改
	db.Raw("select * from information where `name`='李'")                                                 //查
	//http.HandleFunc("/", model.Post)
	//err = http.ListenAndServe(":6666", nil)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
}
