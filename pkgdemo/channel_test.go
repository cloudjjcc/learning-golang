package pkgdemo

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan string, 1)
	ticker := time.NewTicker(1 * time.Second)
	time.AfterFunc(10*time.Second, func() {
		ticker.Stop()
		fmt.Println("ticker stopped")
	})
	for {
		select {
		case ts := <-ticker.C:
			ch <- ts.String()
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func TestCloseChan(t *testing.T) {
	ch := make(chan time.Time, 1)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case v := <-ticker.C:
				ch <- v
			case val, ok := <-ch:
				if !ok {
					fmt.Println("not ok")
					continue
				}
				fmt.Println(val)
			}
		}
	}()
	time.AfterFunc(3*time.Second, func() {
		close(ch)
	})
	select {}
}

func TestSendNilChan(t *testing.T) {
	var ch chan int
	go func() {
		fmt.Println("received:", <-ch)
	}()
	ch <- 1 // will block forever
	fmt.Println("exit")
}

func TestChanExample1(t *testing.T) {
	// send nil chan
	var ch chan int
	//ch <- 1
	select {
	case ch <- 1:
	default:

	}
}
