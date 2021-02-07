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
		case ts, ok := <-ticker.C:
			if !ok {
				return
			}
			ch <- ts.String()
		case msg := <-ch:
			fmt.Println(msg)
		default:
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
