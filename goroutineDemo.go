package main

import (
	"fmt"
	"time"
)

func test()  {
	for i:=0;i<1000;i++  {
		fmt.Println("test=====",i);
		time.Sleep(time.Second)
	}
}

func main(){
	go test();
	for i:=0;i<1000 ;i++  {
		fmt.Println("main=====",i);
		time.Sleep(time.Second)
	}
}
