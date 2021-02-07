package pkgdemo

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("bey bey")
	fmt.Println("do something ...")
}

func TestDefer2(t *testing.T) {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	time.Sleep(1 * time.Second)
}

func TestDefer3(t *testing.T) {
	defer ((func())(nil))()
	fmt.Println("do something ...")
}
