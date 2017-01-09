package main

import (
    "fmt"
)

func main(){
        a()
        b()
        c()
		
		//析构方法
		/*defer func(a int){
		    fmt.Println(a)
		}(12)*/
		
		defer d()
	    fmt.Println(f(1,2,3))
		a, b := swap("Mahesh", "Kumar")
        fmt.Println(a, b)
		
}

func a(){
        fmt.Println("aaa")
}

func b(){
        fmt.Println("bbb")
}

func c(){
        fmt.Println("ccc")
}

func d(){
        fmt.Println("ddd");
}

func f(a,b,c int)(t,r,e int){
        t = a
        r = b
        e = c
		fmt.Println(t)
		fmt.Println(r)
		fmt.Println(e)
		return t,r,e
}
func swap(x, y string) (string, string) {
   return y, x
}