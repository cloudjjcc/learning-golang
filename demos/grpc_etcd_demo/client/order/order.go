package order

import (
	"fmt"
	"sync"
	"time"

	"github.com/cloudjjcc/learning-golang/grpc_etcd_demo/proto"
	clientv3 "go.etcd.io/etcd/client/v3"
	etcdresolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	proto.OrderServiceClient
}

var (
	defaultCli *Client
	mu         sync.Mutex
)

func GetClient(etcdEndpoints []string, serviceName string) (*Client, error) {
	mu.Lock()
	defer mu.Unlock()
	if defaultCli != nil {
		return defaultCli, nil
	}
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: 10 * time.Second,
	})

	builder, err := etcdresolver.NewBuilder(cli)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(serviceName,
		grpc.WithResolvers(builder),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	)
	if err != nil {
		return nil, fmt.Errorf("new grpc conn failed:%w", err)
	}
	client := proto.NewOrderServiceClient(conn)
	defaultCli = &Client{
		OrderServiceClient: client,
	}
	return defaultCli, nil
}
