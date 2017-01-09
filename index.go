package main

import "fmt"
//常量的定义
const (
    PI = 3.14
	const1 = 1
)
//全局变量的声明与
var (
    name string = "xiaoming"
	age int = 321
)

//var a,b,c int = 1,2,3
func main() {
    //局部变量
    /*a,b,c := 1,2,3
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
    
    fmt.Println(name);   
    fmt.Println(age);*/
    //二位数组必须指定数组的长度
	//users := [...][3]int{{1,2},{3,4}}
	
	/*stu := [...]int{1,2,3,4,5,6,7}
	fmt.Println(stu)
	
	length := len(stu)
	
	for i:=0;i<length;i++{
	   fmt.Println(stu[i])
	}*/
	
    //数组
    //a :=[2]int{1,2}
    //fmt.Println(a);	
   
    /*stu := [...]int{1,2,3,4,5,6,7,8,9} 
	fmt.Println(stu[5])
	fmt.Println(stu[5:9])
	fmt.Println(stu[5:])
	fmt.Println(stu[:5])
	fmt.Println(stu[5:len(stu)])*/
	
	//切片
	a := []byte{'a','b','c','d','e','f','g','h','i','m'}
	
	sb := a[2:5]
	
	fmt.Println(string(sb))
	
	fmt.Println(len(a),cap(a))
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
   
   
}
 