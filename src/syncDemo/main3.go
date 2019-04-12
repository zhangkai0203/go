//读写锁  可以是多个读，但是必须是一个写
package main

import (
	"fmt"
	"sync"
	"time"
)

var num int
var mu sync.RWMutex

func main() {

	//写
	go func() {
		for i := 0; i < 10000; i++ {
			mu.Lock()
			num ++
			fmt.Println("读取===", num)
			time.Sleep(time.Second)
			mu.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mu.RLock()
            fmt.Println("读文件1====",num)
			time.Sleep(time.Second)
			mu.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			mu.RLock()
			fmt.Println("读文件2====",num)
			time.Sleep(time.Second)
			mu.RUnlock()
		}
	}()

	time.Sleep(time.Second * 1000)

}
