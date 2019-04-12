//context 用于系统的响应超时
package main

import (
	"context"
	"fmt"
	"time"
)

var ctx context.Context
var b bool = true

//值的存取
func data()  {
	//context 存值
	ctx = context.WithValue(context.Background(),"trace_id",123456)
	ctx = context.WithValue(ctx,"name","zhangsan")

	//context 取值，取出来的值是interface的接口需要转换
	id := ctx.Value("trace_id").(int)
	name := ctx.Value("name").(string)

	fmt.Println(id)
	fmt.Println(name)
}

//第一种
func demo()  {
	//存活一秒钟，过期跳出循环
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	//销毁资源
	defer cancel()

	for  {
		if b == false { return }

		select {
		case <- ctx.Done():
			fmt.Println("过期提示")
			b = false
		default:
			fmt.Println("没有过期")
		}
	}
}

//第二种
func demo1()  {
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second))
	ctx = context.WithValue(ctx,"id",12)

	defer cancel()
	for  {
		if b == false { return }

		select {
		case <- ctx.Done():
			fmt.Println("过期提示===",ctx.Value("id").(int))
			b = false
		default:
			fmt.Println("没有过期===",ctx.Value("id").(int))
		}
	}
}


func main() {
	data()
	//demo()
	//demo1()
}
