package pkgdemo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func counter(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	count := 0
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Println("\ncounter exit")
			return
		case <-ticker.C:
			fmt.Print(count)
			count++
		}
	}
}
func TestContext(t *testing.T) {
	ctx, cancelFn := context.WithCancel(context.Background())
	go counter(ctx)
	stopCh := make(chan struct{})
	time.AfterFunc(10*time.Second, func() {
		cancelFn()
		stopCh <- struct{}{}
	})
	<-stopCh
}
func TestContext2(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	go counter(ctx)
	time.Sleep(10 * time.Second)
}
