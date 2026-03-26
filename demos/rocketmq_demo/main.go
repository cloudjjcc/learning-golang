package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

const (
	demoTopic = "demo"
)

func consume(ctx context.Context) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("test_group"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.88.77:9876"})),
	)
	if err != nil {
		panic(err)
	}
	if err := c.Start(); err != nil {
		panic(err)
	}
	defer func() {
		if err := c.Shutdown(); err != nil {
			fmt.Println("consumer shutdown err:", err)
			return
		}
	}()
	if err = c.Subscribe(demoTopic,
		consumer.MessageSelector{},
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, msg := range ext {
				fmt.Println(msg)
			}
			return consumer.ConsumeSuccess, nil
		}); err != nil {
		fmt.Println("consume err:", err)
	}
	<-ctx.Done()
}
func produce(ctx context.Context) {
	p, err := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"192.168.88.77:9876"})),
	)
	if err != nil {
		panic(err)
	}
	if err := p.Start(); err != nil {
		panic(err)
	}
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
		if err := p.Shutdown(); err != nil {
			fmt.Println("producer shutdown err:", err)
			return
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			msg := primitive.NewMessage(demoTopic, []byte("now is:"+time.Now().String()))
			msg.WithTag("test")
			sendSync, err := p.SendSync(ctx, msg)
			if err != nil {
				fmt.Println("send sync err:", err)
			} else {
				fmt.Println("send sync result:", sendSync)
			}
		}
	}
}
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		consume(ctx)
	}()
	go func() {
		defer wg.Wait()
		produce(ctx)
	}()
	wg.Wait()
}
