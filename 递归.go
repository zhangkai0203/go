package main

import (
	"fmt"
	"time"
)
//递归调用方式
func test(n int) {
	if n == 0 {
		fmt.Println("最后的值=====",n)
		return
	}
	time.Sleep(time.Second)
	fmt.Println(n)
	test(n-1);
}


func main()  {
	test(10)

	func(){
		fmt.Println("aaa")
	}()
}
