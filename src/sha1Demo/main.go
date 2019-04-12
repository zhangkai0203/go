package main

import (
	"crypto/sha1"
	"fmt"
	"sort"
	"strings"
)

func main() {

	m := make(map[string]string)
	var list []string


	m["name"] = "zhangsan"
	m["age"] = "18"
	m["phone"] = "123456"
	m["email"] = "123@163.com"
	m["sing"] = "aaaa"

	for k,v :=  range m {
		list = append(list,fmt.Sprintf("%s=%s",k,v))
	}

	sort.Strings(list)
    lists := strings.Join(list,"&")

	fmt.Println(lists)

    h := sha1.New()
    h.Write([]byte(lists))
    bx := h.Sum(nil)
    fmt.Printf("%x",bx)
}
