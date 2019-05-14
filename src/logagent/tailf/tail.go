package tailf

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

var (
	line *tail.Line
	ok bool
	myChan = make(chan *MsgData,100)
)

type MsgData struct {
	Msg string
	Topic string
}

func InitTail(conf map[string]string)  {

	tails,err := tail.TailFile(conf["collect_path"],tail.Config{
		ReOpen:true,
		Follow:true,
		Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll:true,
	})

	if err != nil {
		fmt.Println("tail file err:%v",err)
        return
	}

	go func() {
		for {
			line,ok = <- tails.Lines
			if !ok {
				fmt.Println("tail file close reopen,filename:%s",tails.Filename)
				time.Sleep(time.Millisecond * 100)
				continue
			}

			if line.Err != nil {
				fmt.Println("tail read file line err:",line.Err)
				continue
			}

			data := &MsgData{line.Text,conf["topic"]}
			myChan <- data
		}
	}()
}

func GetOneLine() (msg *MsgData){
     msg = <- myChan

     fmt.Println(msg)
	 return msg
}


