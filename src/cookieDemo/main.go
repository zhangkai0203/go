package main

import (
	"fmt"
	"net/http"
	"time"
)

func init() {
	//设置时区 上海
	time.LoadLocation("Asia/Shanghai")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index2)
	http.ListenAndServe(":9000", nil)
}

//设置cookie
func index(w http.ResponseWriter, r *http.Request) {
	date := time.Now().AddDate(1, 0, 0)
	cookit := http.Cookie{
		Name:    "gocookieid",
		Value:   "123456789",
		Path:    "/",
		Domain:  "localhost",
		Expires: date,
	}
	http.SetCookie(w, &cookit)
	fmt.Fprint(w, date)
}

func index2(w http.ResponseWriter, r *http.Request) {
	res,_ := r.Cookie("gocookieid")

	fmt.Fprint(w,res)

	for _, c := range r.Cookies() {
		fmt.Fprint(w,c)
	}
}
