package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	m := md5.New()
	m.Write([]byte("123456"))
    bs := m.Sum(nil)
    fmt.Printf("%T==%d==%x",bs,bs,bs)

}
