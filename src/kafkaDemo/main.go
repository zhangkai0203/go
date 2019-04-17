//kafka实例
package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"time"
)

var wg sync.WaitGroup

//生产者
func Producer()  {
	//初始化
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用
	config.Producer.Return.Successes = true

	//生产者
	client, err := sarama.NewSyncProducer([]string{"192.168.1.32:9092"}, config)

	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}

	defer client.Close()

	for i := 1; i < 20; i++ {
		//创建消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = "test"
		msg.Value = sarama.StringEncoder("this is kfa msg . index=%i")

		//发送消息
		pid,offset,err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message err = ",err)
			return
		}
		fmt.Printf("pid ==%v,offset==%v\n",pid,offset)
	}
}

//消费者
func Consumer()  {
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic,msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	consumer.Close()
}

func readData(pc sarama.PartitionConsumer)  {
	for msg := range pc.Messages(){
		fmt.Printf("topic=%v,value=%s\n",msg.Topic,msg.Value)
		continue
	}
}

func main() {
	go Consumer()
	time.Sleep(time.Second * 100)

}
