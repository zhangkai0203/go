//读取带缓冲的文件

package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main()  {
       file,err := os.Open("D:/go/1.txt");
	if err != nil {
		fmt.Println("file errror",err)
	}

	rdader:= bufio.NewReader(file);


	for  {
		str,err := rdader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != err {
			fmt.Println("bufio error",err)
		}
		fmt.Println(str)
	}




}
