//链式调用
package main

import "fmt"

type studen struct {
	Name string
	Age  uint8
}

type stu struct {
	Name string
	Age uint8
}

func (t *studen) setName(name string) {
	t.Name = name
}

func (t *studen) setAge(age uint8) {
    t.Age = age
}
func (t *studen) pr()  {
	fmt.Printf("name:%s,age:%d\n",t.Name,t.Age)
}



func (s *stu) setStuName(name string) *stu {
	s.Name = name
	return s
}

func (s *stu) setStuAge(age uint8) *stu {
    s.Age = age
	return s
}

func (s *stu) stuPr()  {
	fmt.Printf("name=%s,age=%d\n",s.Name,s.Age)
}



func main() {
	//非链式调用
    s := &studen{}
    s.setName("zhangsan")
    s.setAge(12)
    s.pr()

    //链式调用
    st := &stu{}
    st.setStuName("lisi").setStuAge(16).stuPr()

}
