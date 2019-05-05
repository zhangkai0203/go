package main

import (
	"fmt"
	"os"
)

func main()  {
	file,err := os.Create("demo.txt")
	if err != nil {
		fmt.Println("创建文件错误")
		return
	}
	ret,err := file.WriteString("hello world")
	if err != nil {
		fmt.Println("写入信息错误")
	}
	fmt.Println(ret)
}
