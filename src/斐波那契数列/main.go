package main

import "fmt"

var (
	a, b int = 1, 1
)

//斐波那契数列
func main() {
	for i := 0; i < 10; i++ {
		temp := a
		a = b
		b = a + temp
		fmt.Println(b)
	}
}
