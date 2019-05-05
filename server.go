package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
)

func index(w http.ResponseWriter,r *http.Request)  {
     io.WriteString(w,"index\n");


         fmt.Println(r)
	//fmt.Printf("%v\n",r)
}

func goods(w http.ResponseWriter,r *http.Request)  {
     io.WriteString(w,"goods")
}

func main()  {
    http.HandleFunc("/",index);
    http.HandleFunc("/goods-*",goods);
    err := http.ListenAndServe(":8000",nil);
    if err != nil{
	    log.Fatal("server",err)
    }
}
