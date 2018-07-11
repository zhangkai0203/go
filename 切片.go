package main

import "fmt"

func main() {
	//定义
	var data []int
	var data2 []int = make([]int,10)
	data1 := make([]int,10)
	fmt.Println(data,data1,data2)

	//初始化
	rs := []int{1,2,3,4,5,6,7,8,9}
	for i:=0;i < len(rs) ; i++ {
		data1[i] = rs[i]
	}
	data = append(data,1,2,3,4,5,6,7,8,9)
	fmt.Println(rs,data1,data)

	fmt.Println(data[:])
	fmt.Println(data[2:])
	fmt.Println(data[2:5])

}
