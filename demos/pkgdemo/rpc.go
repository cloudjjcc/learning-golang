package pkgdemo

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type api struct {
}

func (apip *api) Hello(req string, resp *string) error {
	*resp = "reply:" + req
	return nil
}
func serve() {
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	server := rpc.NewServer()
	if err := server.RegisterName("api", &api{}); err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func client() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		panic(err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	reply := ""
	if err := client.Call("api.Hello", "client", &reply); err != nil {
		log.Fatal(err)
		return
	}
	log.Println(reply)
}
