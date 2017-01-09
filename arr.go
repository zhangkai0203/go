package main

import "fmt"
var sum int
func main(){
        //声明一个数组
		arr := []int{1,2,3,4,5,6}
		
		var a int = 1
		
		fmt.Println(arr[5])
		fmt.Println(arr[3:])
		fmt.Println(arr[:3])
		//求和
		s(arr)
		fmt.Println(sum)
		
		//指针
		strwor(&a)
		fmt.Println(a)
}
func s(a []int){
 
    for i:=0;i<len(a);i++{
	    sum += a[i]		
	}
}
func strwor(a *int){
    *a = 12
}