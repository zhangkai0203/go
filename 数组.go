package main

import "fmt"

func main() {
	//定义
	var arr [10]int
	fmt.Println(arr)
        //初始化值
	arr =  [10]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(arr)

	for key,val:= range arr{
		arr[key] = val * 10
	}
	fmt.Println(arr)

	arr[9] = 100

	fmt.Println(arr)
}
