package main

import (
	"fmt"
	"crypto/md5"
)

func md5Demo(m string) string{
	md := md5.Sum([]byte(m))
	data := fmt.Sprintf("%x",md)
	return data
}



func test()  {
	var name string
	var age int

	fmt.Scanln(&name,&age)
	fmt.Println("name:",name,"======","age",age)
}

func test2()  {
	pass := md5.Sum([]byte("123456"))
	pa := fmt.Sprintf("%x",pass)
	fmt.Println(pa)
}

func test3()  {
	md := md5.New()
	md.Write([]byte("123456"))
	pa := md.Sum(nil)
	data := fmt.Sprintf("%x",pa)
	fmt.Println(data)
}


func main()  {
     //test2()
	//test3()
	pass := md5Demo("123456")
	fmt.Println(pass)
}
