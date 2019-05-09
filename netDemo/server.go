package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("ln error", err)
		return
	}
	for {
		conn, err := ln.Accept();
		if err != nil {
			fmt.Println("conn error", err)
			return;
		}
		fmt.Println(conn)
	}
}
