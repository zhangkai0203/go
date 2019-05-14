package main

import (
	"io"
	"net/http"
	"session"
)

func login(w http.ResponseWriter,r *http.Request)  {
    session.Set(w)

}

func index(w http.ResponseWriter,r *http.Request)  {
	session.Get(w,r)
	io.WriteString(w,"hello world")
}

func logou(w http.ResponseWriter,r *http.Request)  {
    session.Delete(w,r)
}

func main() {

    http.HandleFunc("/",index)
    http.HandleFunc("/login",login)
    http.HandleFunc("/logou",logou)
    http.ListenAndServe(":9000",nil)
}
