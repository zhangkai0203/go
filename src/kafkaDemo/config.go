package main

import (
	"fmt"
	"src/github.com/astaxie/beego/config"
)

func main() {

    conf,err := config.NewConfig("ini","config.conf")
	if err != nil {
		fmt.Println("new config file err:",err)
		return
	}

	port,err := conf.Int("server::port")
	if err != nil {
		fmt.Println("port err:",err)
		return
	}
	fmt.Println("port data:",port)


}
