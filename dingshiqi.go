package main
//go语言定时器功能
import (
	"time"
	"fmt"
)

func test()  {
	var num int = 0
	//每秒请求一次
	t := time.NewTicker(time.Second);
	for v := range t.C {
		if num == 10 {
			t.Stop()
			break;
		}
		fmt.Println(v,num)
		num++
	}
}

func test2()  {
        t := time.NewTimer(time.Second * 2);
	v := <- t.C
	fmt.Println(v)
}



func main()  {
	test()
	//test2()

}