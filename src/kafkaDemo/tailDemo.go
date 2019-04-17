package main

import (
	"fmt"
	"src/github.com/hpcloud/tail"
	"time"
)

func main() {

	filename := "./my.log"

	tails,err := tail.TailFile(filename,tail.Config{
		ReOpen:true,
		Follow:true,
		//Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll:true,

	})

	if err != nil {
		fmt.Println("tail file err:",err)
		return
	}

	var msg *tail.Line
	var ok bool

	for {
		msg,ok = <- tails.Lines

		if !ok {
			fmt.Println("tail file close reopen,filename:%s\n",tails.Filename)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		fmt.Println("msg:",msg)
	}


}
