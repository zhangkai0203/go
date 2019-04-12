//静态服务器
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type body struct {
	Title string
}

func index(w http.ResponseWriter,r *http.Request)  {
	t,err := template.ParseFiles("view/layout.html","view/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w,body{"主体信息"})
}

func main() {

	//启动静态服务
	h := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", h)) // 启动静态文件服务

    http.HandleFunc("/",index)
	http.ListenAndServe(":9000",nil)
}
