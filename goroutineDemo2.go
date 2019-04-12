package main

import (
	"fmt"
	"time"
)

func main()  {
	start := time.Now().Unix();
	for i:=1;i<1000000 ;i++  {
		go pr(i)
	}
	time.Sleep(time.Second)

	end := (time.Now().Unix()) - start

	fmt.Println("时间：",end)
}

func pr(i int)  {
	fmt.Println("并发模型：",i)
}
