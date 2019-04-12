//加锁的情况  互斥锁
package main

import (
	"fmt"
	"sync"
	"time"
)

var num int
var mu sync.Mutex

func main() {
	go func() {
		for i:=0;i< 10000 ;i++  {
			mu.Lock()
			num ++
			mu.Unlock()
		}
	}()

	go func() {
		for i:=0;i< 10000 ;i++  {
			mu.Lock()
			num ++
			mu.Unlock()
		}
	}()

	go func() {
		for i:=0;i< 10000 ;i++  {
			mu.Lock()
			num ++
			mu.Unlock()
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("结果==",num)

}