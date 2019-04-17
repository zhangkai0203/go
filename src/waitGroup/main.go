//等待gorouting 结束后再结束主进程实例
package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup


func main() {

	for i:=1;i<10 ;i++  {
		w.Add(1)
		go readData(i)
	}

	w.Wait()
	fmt.Println("结束程序")
}

func readData(i int)  {
	fmt.Println("num====",i)
	time.Sleep(time.Second)
	w.Done()
}




