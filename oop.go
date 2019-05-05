package main

import "fmt"

type mianji struct {
	width int
	height int
}

func (s mianji) fangzi() int {
	return s.width * s.height
}

func main()  {
	mianji := mianji{width:100,height:50}

	fmt.Println(mianji.fangzi())

        mianji.width = 10
	mianji.height= 5

	fmt.Println(mianji.fangzi())
}
