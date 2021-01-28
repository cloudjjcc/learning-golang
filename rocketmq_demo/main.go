package main

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/core"
	"log"
	"sync"
	"time"
)

// 推送消费者demo
func pushConsumerDemo(group *sync.WaitGroup) {
	defer group.Done()
	config := &rocketmq.PushConsumerConfig{
		ClientConfig: rocketmq.ClientConfig{
			GroupID:    "demo_consumer",
			NameServer: "127.0.0.1:9876",
		},
		Model:         rocketmq.Clustering,
		ConsumerModel: rocketmq.Orderly,
	}
	consumer, err := rocketmq.NewPushConsumer(config)
	if err != nil {
		log.Fatalln(err)
	}
	err = consumer.Subscribe("demoTopic", "*", func(msg *rocketmq.MessageExt) rocketmq.ConsumeStatus {
		log.Println(msg.Body)
		return rocketmq.ConsumeSuccess
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = consumer.Start()
	if err != nil {
		log.Fatalln(err)
	}
	time.AfterFunc(time.Minute*1, func() {
		defer consumer.Shutdown()
		log.Println("consumer shutdown....")
	})
}

// 生产者demo
func producerDemo(group *sync.WaitGroup) {
	defer group.Done()
	config := &rocketmq.ProducerConfig{
		ClientConfig: rocketmq.ClientConfig{
			GroupID:    "demo_producer",
			NameServer: "127.0.0.1:9876",
		},
		ProducerModel: rocketmq.CommonProducer,
	}
	producer, err := rocketmq.NewProducer(config)
	if err != nil {
		log.Fatalln(err)
	}
	err = producer.Start()
	if err != nil {
		log.Fatalln(err)
	}
	defer producer.Shutdown()
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case s := <-ticker.C:
			result, err2 := producer.SendMessageSync(&rocketmq.Message{
				Topic: "demoTopic",
				Body:  fmt.Sprintf("ticker %s", s.String()),
			})
			if err2 != nil {
				log.Println(err2)
			} else {
				log.Println(result)
			}
		default:
		}
	}
}
func main() {
	log.Println("start test rocket mq...")
	group := &sync.WaitGroup{}
	group.Add(2)
	go pushConsumerDemo(group)
	go producerDemo(group)
	(*group).Wait()
	log.Println("exit...")
}
