package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"logagent/tailf"
)

func ServerRun()  {

	for  {
		msg := tailf.GetOneLine()
		fmt.Println(msg)
		fmt.Println("serverRun..........")
		sendTokafka(msg)
	}
}

func sendTokafka(m *tailf.MsgData)  {

	//初始化
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用
	config.Producer.Return.Successes = true

	//生产者
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}

	defer client.Close()

	//创建消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = m.Topic
	msg.Value = sarama.StringEncoder(m.Msg)

	//发送消息
	pid,offset,err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message err = ",err)
		return
	}
	fmt.Printf("pid ==%v,offset==%v\n",pid,offset)
}

