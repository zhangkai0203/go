package main

import (
    "fmt"
)

func main(){
    m1 := map[int]string{1:"a",2:"b",3:"c"}
	
	fmt.Println(m1)
	
	//声明m2
	m2 := make(map[string]int)
	
	for k,v := range m1{
	    m2[v] = k
	}
	
	fmt.Println(m2)
	
	fmt.Println(a())
	
	b(1,2,3,4)
	
}
//go函数

func a()(a,b,c int){
    a,b,c = 1,2,3
	return a,b,c
}

func b(a ...int){
    fmt.Println(a[0])
	a[0] = 9
	fmt.Println(a)
	
    //fmt.Println(b)
    //fmt.Println(c)
    //fmt.Println(d)
}

