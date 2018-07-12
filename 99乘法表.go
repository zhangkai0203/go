package main

import "fmt"

func main() {
	//正序乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println()
	//倒序
	for z:=9;z>=1 ;z--  {
		for y:=9;y>=z ;y--  {
			fmt.Printf("%d * %d = %d\t",z,y,z*y)
		}
		fmt.Println()
	}
	fmt.Println()
	for o := 9; o >= 1; o-- {
		for p := 1; p <= o; p++ {
			fmt.Printf("%d * %d = %d\t", o, p, o*p)
		}
		fmt.Println()
	}


}
