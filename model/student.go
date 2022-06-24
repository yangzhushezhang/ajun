package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Information struct {
	ID            uint
	Name          string
	Sex           string
	Age           int
	CreationTime  string
	ModifyTheTime string
}

type POST struct {
	Name string
	Sex  string
	Age  int
}

//插入数据
func Post(w http.ResponseWriter, r *http.Request) {
	var A []string
	r.ParseForm()
	fmt.Println(r.Form)
	A = append(A, r.Form["name"]...)
	A = append(A, r.Form["sex"]...)
	A = append(A, r.Form["age"]...)

	B := A[0]
	C := A[1]
	D := A[2]

	dsn := "root:123456@tcp(127.0.0.1:3306)/test0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("错误:", err.Error())
	}
	a := make([]Information, 0)
	db.Raw("INSERT INTO information(name,sex,age) VALUES  ('" + B + "','" + C + "'," + D + ")").Scan(&a)
	fmt.Println(A)
	fmt.Fprintf(w, "已入库")
}

//查询表格信息
func check(A string) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("错误:", err.Error())
		a := make([]Information, 0)
		db.Raw("SELECT *FROM " + A).Scan(&a)
	}
}
