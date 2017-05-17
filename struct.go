package main

import "fmt"

//构造方法
/*type person struct{
    name string
	age  int
}*/

//嵌套
/*type person struct{
    name string
	age  int
	contact struct{
	    phone,city string
	}
}*/

//继承
type xuexiao struct{
    sex int
}

type a struct{
    xuexiao
    name string
	age  int
}
type b struct{
    xuexiao
    name string
	age  int
}

func main(){
    //第一种默认值的方法
	/*a := person{}
	a.name = "zhangsan"
	a.age = 14*/
	
	//第二种默认值的方法
	/*a := person{
	    name : "zhangsan",
		age  : 19,
	}*/
	
	//第三种默认值方法
	/*a := struct{
        name string
        age  int		
	}{
	    name : "zhangsan",
		age  : 19,
	}*/
	
	/*a := person{name:"zhangsan",age:19}
	a.contact.phone = "123"
	a.contact.city  = "beijing"*/
	
	a := a{name:"zhangsan",age:17,xuexiao:xuexiao{sex:0}}
	b := b{name:"lisi",age:19,xuexiao:xuexiao{sex:1}}
	
	fmt.Println(a)
	fmt.Println(b)
	
	
	
	package main

	import "fmt"

	type jiegou struct{
	    width, height int
	}


	func main() {
	    sum := jiegou{width:100,height:80}

	    mj := sum.mianji()

	    fmt.Println(mj)
	}

	func (r jiegou) mianji() int{
	    return r.width * r.height
	}

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
    
}
