package order

import (
	"context"
	"fmt"

	"github.com/cloudjjcc/learning-golang/grpc_etcd_demo/proto"
)

var _ proto.OrderServiceServer = (*Service)(nil)

type Service struct {
	proto.UnimplementedOrderServiceServer
	Addr string
}

func (srv *Service) CreateOrder(ctx context.Context, req *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	fmt.Println("create order:", req.String())
	return &proto.CreateOrderResp{
		Result: "handle by:" + srv.Addr,
	}, nil
}
