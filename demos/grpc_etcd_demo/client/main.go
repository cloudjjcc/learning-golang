package main

import (
	"context"
	"fmt"

	"github.com/cloudjjcc/go-exercises/grpc_etcd_demo/client/order"
	"github.com/cloudjjcc/go-exercises/grpc_etcd_demo/proto"
)

func main() {
	client, err := order.GetClient([]string{
		"192.168.88.77:2379",
		"192.168.88.77:2381",
		"192.168.88.77:2383",
	}, "etcd:///services/order")
	if err != nil {
		panic("get client failed:" + err.Error())
	}
	resp, err := client.CreateOrder(context.Background(), &proto.CreateOrderReq{
		Item:  "test1",
		Count: 10,
	})
	if err != nil {
		fmt.Println("create order failed:" + err.Error())
		return
	}
	fmt.Println("got resp:" + resp.String())
}
