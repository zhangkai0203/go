package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"flag"
)

var (
	pool *redis.Pool
	redisServer  = flag.String("tcp","localhost:6379","");
)

/**
   redis 连接池
   redis 下载地址 https://github.com/gomodule/redigo
   redis 连接池 文档https://studygolang.com/articles/12230
 */

func newPool(addr string) *redis.Pool  {
	return &redis.Pool{
		MaxIdle:     8,  //最大空闲连接数
		MaxActive:   0,   //表示和数据的最大链接数  0：不限制
		IdleTimeout: 300, //最大空闲时间
		Dial: func() (redis.Conn, error) {//初始化链接
			return redis.Dial("tcp", addr)
		},
	}
}

func main() {
        flag.Parse();
	pool = newPool(*redisServer);

	conn := pool.Get();  //从链接池中取出一个连接
	defer conn.Close();  //关闭连接池，关闭不能从连接池中取出链接

	//添加
 	_,err := conn.Do("set","username","zhaoliuaa")
	if err != nil {
		fmt.Println("redis set  error:",err)
	}
        //取出
	name,err := redis.String(conn.Do("get","username"));
	fmt.Println(name)

}
