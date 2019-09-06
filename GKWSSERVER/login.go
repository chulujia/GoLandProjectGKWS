package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)


func loginFunc(respose http.ResponseWriter, request *http.Request) {
	fmt.Println("method:", request.Method) //获取请求的方法
	if request.Method == "GET" {
		tmp, _ := template.New("login.html").ParseFiles("./html/login.html")
		tmp.Execute(respose, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", request.FormValue("username"))
		fmt.Println("password:", request.FormValue("password"))
	}
}

func main() {
	http.HandleFunc("/login", loginFunc)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/*
http://localhost:9090/login
http://192.168.219.134:9090/login
*/

/*
HTML 教程:
https://www.w3school.com.cn/html/index.asp

Go Web编程:
https://www.w3cschool.cn/yqbmht/7rcvwcqm.html
*/