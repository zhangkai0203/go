package main

import (
     "fmt"
     "logagent/config"
     "logagent/kafka"
     "logagent/log"
     "logagent/tailf"
)


func main() {

     //读取配置文件
     conf,err := config.LoadConf()
     if err != nil {
          fmt.Println("读取配置错误",err)
          return
     }
     //启动日志
     log.InitLogger(conf)
     //读取文件
     tailf.InitTail(conf)
     //启动服务
     kafka.ServerRun()
}
