//为什么加锁 一个小例子,输出的结果不一定是想要的
package main

import (
	"fmt"
	"time"
)

var num int

func main() {

	go func() {
		for i:=0;i< 10000 ;i++  {
            num ++
		}
	}()

	go func() {
		for i:=0;i< 10000 ;i++  {
			num ++
		}
	}()

	go func() {
		for i:=0;i< 10000 ;i++  {
			num ++
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("结果===",num)

}
