package main

import "fmt"

func main(){
        var a float32 = 10.8
		var b float32 = 20.6 
		var c float32 = 0
		
		c = sum(a,b)
		fmt.Println(c)
}
func sum(a,b float32)float32{
    return a+b;
}