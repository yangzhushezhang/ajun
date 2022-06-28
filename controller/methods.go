package controller

import (
	"ajun/model"
	_ "ajun/model"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("错误:", err.Error())
	}
	a := make([]model.Information, 0)
	var A []string
	r.ParseForm()

	A = append(A, r.Form["ID"]...)
	A = append(A, r.Form["name"]...)
	A = append(A, r.Form["sex"]...)
	A = append(A, r.Form["age"]...)
	A = append(A, r.Form["action"]...)

	B := A[0]
	C := A[1]
	D := A[2]
	E := A[3]
	F := A[4]

	if F == "insert" && F != "" { //增
		fmt.Println(D)
		db.Raw("INSERT INTO information(name,sex,age) VALUES  ('" + C + "','" + D + "'," + E + ")").Scan(&a)
		fmt.Fprintf(w, "已入库")
		return
	}

	if F == "delete" && F != "" { //删
		db.Raw("DELETE FROM information WHERE id='" + B + "'").Scan(&a)
		fmt.Fprintf(w, "已清除")
		return
	}

	if F == "change" && F != "" { //改
		db.Raw("UPDATE information SET ID='" + B + "',`name`='" + C + "',sex=" + D + ",age=25 WHERE  `name`='" + B + "'").Scan(&a)
		fmt.Fprintf(w, "已修改")
		return
	}

	if F == "check" && F != "" { //查
		db.Raw("select * from information where id='" + B + "'").Scan(&a)
		q, err := json.Marshal(a)
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		w.Write(q)
		return
	} else if F == "checkq" && F != "" {
		db.Raw("select * from information").Scan(&a)
		q, err := json.Marshal(a)
		if err != nil {
			fmt.Println("err:", err.Error())
		}
		w.Write(q)
		return
	}

	fmt.Println("未检测到")
	fmt.Fprintf(w, "action不能为空")
}
