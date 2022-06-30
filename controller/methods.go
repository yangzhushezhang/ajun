package controller

import (
	"ajun/dao/mysql"
	"ajun/model"
	_ "ajun/model"
	"ajun/tools"
	"net/http"
	"strconv"
)

//func Post(w http.ResponseWriter, r *http.Request) {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test0?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open("Mysql", &gorm.Config{})
//	if err != nil {
//		fmt.Println("错误:", err.Error())
//	}
//	a := make([]model.Information, 0)
//	var A []string
//	r.ParseForm()
//
//	A = append(A, r.Form["ID"]...)
//	A = append(A, r.Form["name"]...)
//	A = append(A, r.Form["sex"]...)
//	A = append(A, r.Form["age"]...)
//	//A = append(A, r.Form["form"]...)
//	A = append(A, r.Form["action"]...)
//
//	B := A[0]
//	C := A[1]
//	D := A[2]
//	E := A[3]
//	F := A[4]
//	G := "information"
//
//	if F == "insert" && F != "" { //增
//		fmt.Println(D)
//		db.Raw("INSERT INTO " + G + "(name,sex,age) VALUES  ('" + C + "','" + D + "'," + E + ")").Scan(&a)
//		fmt.Fprintf(w, "已入库")
//		return
//	}
//
//	if F == "delete" && F != "" { //删
//		db.Raw("DELETE FROM " + G + " WHERE id='" + B + "'").Scan(&a)
//		fmt.Fprintf(w, "已清除")
//		return
//	}
//
//	if F == "change" && F != "" { //改
//		db.Raw("UPDATE " + G + " SET ID='" + B + "',`name`='" + C + "',sex='" + D + "',age=" + E + " WHERE  `id`='" + B + "'").Scan(&a)
//		fmt.Fprintf(w, "已修改")
//		return
//	}
//
//	if F == "check" && F != "" { //查
//		db.Raw("select * from " + G + " where id='" + B + "'").Scan(&a)
//		q, err := json.Marshal(a)
//		if err != nil {
//			fmt.Println("err:", err.Error())
//		}
//		w.Write(q)
//		return
//	} else if F == "checkq" && F != "" {
//		db.Raw("select * from " + G).Scan(&a)
//		q, err := json.Marshal(a)
//		if err != nil {
//			fmt.Println("err:", err.Error())
//		}
//		w.Write(q)
//		return
//	}
//
//	fmt.Println("未检测到")
//	fmt.Fprintf(w, "action不能为空")
//}

func Test(w http.ResponseWriter, r *http.Request) {
	//插入
	action := r.FormValue("action")
	if action == "add" {
		atom, _ := strconv.Atoi(r.FormValue("age"))
		add := model.Information{Name: r.FormValue("name"), Sex: r.FormValue("sex"), Age: atom}
		if add.Add(mysql.DB) == false {
			tools.ReturnJson(w, -101, nil, "不要重复添加")
			return
		}
		tools.ReturnJson(w, 200, nil, "添加成功")
		return
	}
	//改
	if action == "update" {
		atom, _ := strconv.Atoi(r.FormValue("age"))
		aaa, _ := strconv.Atoi(r.FormValue("id"))
		adb := model.Information{ID: uint(aaa), Name: r.FormValue("name"), Sex: r.FormValue("sex"), Age: atom}
		if adb.Adb(mysql.DB) == false {
			tools.ReturnJson(w, -101, nil, "修改失败")
			return
		}
		tools.ReturnJson(w, 200, nil, "修改成功")
		return
	}
	//查
	if action == "check" {
		aaa, _ := strconv.Atoi(r.FormValue("id"))
		//adb := model.Information{ID: uint(aaa)}
		//
		//if adb.Adc(mysql.DB) == false {
		//	return
		//}

		in := model.Information{}

		mysql.DB.Where("id=?", aaa).First(&in)

		tools.ReturnJson(w, 200, in, "ok")

		return
	}

	//删
	if action == "delete" {
		bbb, _ := strconv.Atoi(r.FormValue("id"))

		aba := model.Information{ID: uint(bbb)}
		if aba.DEL(mysql.DB) == false {
			tools.ReturnJson(w, -101, nil, "删除失败")
			return
		}
		tools.ReturnJson(w, 200, nil, "删除成功")
		return
	}

	tools.ReturnJson(w, -101, nil, "对不起,请先填写action参数")
}
