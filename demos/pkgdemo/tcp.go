package pkgdemo

import (
	"io"
	"log"
	"net"
)

func echoServer() {
	listener, err := net.Listen("tcp", ":7777")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err:%s", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			for {
				io.Copy(conn, conn)
			}
		}(conn)
	}
}

func send(data string) error {
	conn, err := net.Dial("tcp", "localhost:7777")
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(data))
	if err != nil {
		return err
	}
	buffer := make([]byte, 1000)
	n, err := conn.Read(buffer)
	if err != nil {
		return err
	}
	conn.Close()
	log.Println(string(buffer[:n]))
	return err
}
