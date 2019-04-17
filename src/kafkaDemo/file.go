package main

import (
	"fmt"
	"os"
)

func main() {


	filename := "./my.log"

	for i:=1;i<1000 ;i++  {

		str := fmt.Sprintf("hello world file %d\n",i)

		file,err := os.OpenFile(filename,os.O_RDWR|os.O_APPEND,0777)
		if err != nil {
			fmt.Println("open file err:",err)
			return
		}

		n,err := file.Write([]byte(str))
		if err != nil {
			fmt.Println("write err:",err)
			return
		}
		fmt.Println("data====",n)
	}





}



