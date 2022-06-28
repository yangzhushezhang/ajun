package main

import (
	"ajun/controller"
	_ "ajun/model"
	"fmt"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.Post)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
}
