package main

import "fmt"

var myMap map[string]string

type Stu interface {
	Get() string
}

func getData(list interface{})  {

    s := list.(User)
	fmt.Println(s.Name)
}

type User struct {
	Name string
}

func (u User) Get() string{
	return u.Name
}

func main() {

	var stu Stu
	stu = User{"zhangsan"}
	fmt.Println(stu.Get())

	user := User{"lisi"}

	var s Stu
	s = user
	fmt.Println(s.Get())

	/*myMap = make(map[string]string)
    myMap["name"] = "zhangsan"
    myMap["age"] = "13"*/

   u :=  User{"wangwu"}


	getData(u)


}