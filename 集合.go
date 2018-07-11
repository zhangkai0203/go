package main

import "fmt"

func main() {
	//定义
	var data =  map[string]string{};
	data1 := make(map[string]string)
	fmt.Println(data,data1)

	//初始化
	re := map[string]string{"name":"zhangsan"}
	data1["name"] = "zhangsan"
	fmt.Println(re,data1)

	//遍历
	for key,val:= range re{
		fmt.Println(key,"===",val)
	}

	//添加
	re["age"] = "12"
	fmt.Println(re)

	//删除
	delete(re,"age")
	fmt.Println(re)


}
