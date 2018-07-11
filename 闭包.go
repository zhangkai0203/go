package main

import "fmt"
//不带参数
func test() func() int {
	s := 0
	return func() int {
		s += 1
		return s
	}
}
//带参数
func test1(n, y int) func() int {
	s := 0;
	return func() int {
		s = n + y + s
		return s
	}
}
func test2() func(n int) int {
        var s int = 0
	return func (d int) int{
	     s += d
		return s
	}
}
func test3(n int) func () int {
	s := 0
	return func() int {
		s = n * 10
		n++
	       return s
	}

}
func main() {

	/*f := test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())*/

	/*z := test1(1,2)
	fmt.Println(z())
	fmt.Println(z())
	fmt.Println(z())
	fmt.Println(z())*/

	/*y := test2()
	fmt.Println(y(1))
	fmt.Println(y(1))
	fmt.Println(y(1))
	fmt.Println(y(1))*/

	m := test3(1)
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())





}
