package register

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type EtcdRegister struct {
	cli           *clientv3.Client
	ttl           int64
	leaseID       clientv3.LeaseID
	servicePrefix string
	em            endpoints.Manager
	ctx           context.Context
	cancel        context.CancelFunc
}

func NewEtcdRegister(ctx context.Context, etcdEndpoints []string, servicePrefix string) (*EtcdRegister, error) {
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: etcdEndpoints,
	})
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	r := &EtcdRegister{
		ctx:           ctx,
		cancel:        cancel,
		cli:           etcdCli,
		ttl:           30,
		servicePrefix: servicePrefix,
	}
	r.em, err = endpoints.NewManager(r.cli, servicePrefix)
	if err != nil {
		return nil, err
	}
	if err := r.init(); err != nil {
		return nil, err
	}
	return r, nil
}
func (r *EtcdRegister) init() error {
	// 租约，服务挂掉自动删除
	lease, err := r.cli.Grant(r.ctx, r.ttl)
	if err != nil {
		return err
	}
	r.leaseID = lease.ID
	alive, err := r.cli.KeepAlive(r.ctx, r.leaseID)
	if err != nil {
		return err
	}
	go func() {
		defer func() {
			fmt.Println("keep alive exit")
		}()
		for resp := range alive {
			fmt.Println(resp)
		}
	}()
	return nil
}

// RegisterEndpoint 注册一个 endpoint
func (r *EtcdRegister) RegisterEndpoint(serviceName, addr string) error {
	return r.em.AddEndpoint(r.ctx, fmt.Sprintf("%s/%s/%s", r.servicePrefix, serviceName, addr), endpoints.Endpoint{Addr: addr}, clientv3.WithLease(r.leaseID))
}

func (r *EtcdRegister) Close() {
	r.cancel()
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	if _, err := r.cli.Revoke(ctx, r.leaseID); err != nil {
		fmt.Println(err)
	}
	if err := r.cli.Close(); err != nil {
		fmt.Println(err)
	}
}
