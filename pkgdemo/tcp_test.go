package pkgdemo

import (
	"testing"
	"time"
)

func TestEchoServer(t *testing.T) {
	go echoServer()
	// wait server start
	time.Sleep(5 * time.Second)
	send("hello world")
	//send("你好，世界")
}
