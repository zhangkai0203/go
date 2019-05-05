package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server....");
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("server error ", err)
		return;
	}
	for {
		conn, err := ln.Accept();
		if err != nil {
			fmt.Println("accept error ", err)
			continue;
		}
		go process(conn)
	}

}
func process(conn net.Conn) {
	defer conn.Close();
	for {
		buf := make([]byte, 512)
		_,err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error ",err)
			return
		}
		fmt.Println("read",string(buf))
	}
}
