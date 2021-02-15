package pkgdemo

import (
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {
	go runHttpServer()
	time.Sleep(5 * time.Second)
	clientCall()
}
