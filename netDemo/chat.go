package main

import (
	"net"
	"fmt"
	"bufio"
)

func main()  {
	conn,err := net.Dial("tcp","localhost:9000")
	if err != err {
		fmt.Println("conn error = ",err)
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status);
}
