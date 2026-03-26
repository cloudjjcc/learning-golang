package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"sync"
	"syscall"

	"github.com/cloudjjcc/go-exercises/grpc_etcd_demo/proto"
	"github.com/cloudjjcc/go-exercises/grpc_etcd_demo/register"
	"github.com/cloudjjcc/go-exercises/grpc_etcd_demo/server/order"
	"google.golang.org/grpc"
)

func main() {
	serviceName := "order"
	addr := "127.0.0.1:9000"

	ctx := context.Background()
	var stop context.CancelFunc
	ctx, stop = signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// 初始化注册器
	etcdRegister, err := register.NewEtcdRegister(ctx, []string{
		"192.168.88.77:2379",
		"192.168.88.77:2381",
		"192.168.88.77:2383",
	}, "services")
	if err != nil {
		panic("create new etcd register error:" + err.Error())
	}

	// 创建新的grpc服务
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("listen error:" + err.Error())
	}
	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, &order.Service{Addr: addr})
	wg := sync.WaitGroup{}
	wg.Add(1)

	// 启动grpc服务
	go func() {
		defer func() {
			wg.Done()
		}()
		log.Println("Server running at:", addr)
		if err := grpcServer.Serve(lis); err != nil {
			fmt.Println(err)
		}
	}()

	// 注册服务
	if err := etcdRegister.RegisterEndpoint(serviceName, addr); err != nil {
		panic("register service error:" + err.Error())
	}

	<-ctx.Done()
	etcdRegister.Close()
	grpcServer.GracefulStop()
	wg.Wait()
}
