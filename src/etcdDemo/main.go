package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	cli *clientv3.Client
	key string = "nginx"
)

type logConf struct {
	Path string   `json:"path"`
	Topic string  `json:"topic"`
} 

func init()  {
	cl,err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("etcd error==",err)
		return
	}
	cli = cl
}


func insert()  {
	//设置1秒超时，访问etcd有超时控制
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	_,err := cli.Put(ctx,key,"lisi")
	//操作完毕取消etcd
	cancel()
	if err != nil {
		fmt.Println("etcd put error==",err)
		return
	}
}

func insert1()  {
	//设置1秒超时，访问etcd有超时控制
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	
	var logConfArr []logConf
    logConfArr = append(
    	logConfArr,
    	logConf{
    		"/nginx/error.log",
    		"nginx",
		})
	logConfArr = append(
		logConfArr,
		logConf{
			"/nginx/success.log",
			"nginx",
		})

	data,_ := json.Marshal(logConfArr)
	
	_,err := cli.Put(ctx,key,string(data))
	//操作完毕取消etcd
	cancel()
	if err != nil {
		fmt.Println("etcd put error==",err)
		return
	}
}

func get()  {
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	resp,err := cli.Get(ctx,key)
	cancel()
	if err != nil {
		fmt.Println("etcd get error",err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func main() {
    //insert()
    insert1()
    get()
	defer cli.Close()

}
