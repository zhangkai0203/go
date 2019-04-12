package main

import (
	"fmt"
	"time"
)

var m1 chan int

func writeData(m1 chan int) {
	for i := 0; i < 100000; i++ {
		m1 <- i;
		fmt.Println("write")
	}
}

func readData(m1 chan int) {
	for i := 0; i < 100000; i++ {
		n := <-m1
		fmt.Println(n)
	}
}

func main() {
	m1 = make(chan int, 100000);
	go writeData(m1)
	go readData(m1)
	time.Sleep(time.Second)
}
