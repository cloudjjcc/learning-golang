package leetcode

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func sequencePrint(n int) {
	chs := make([]chan struct{}, n)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	exitCh := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		chs[i] = make(chan struct{})
		wg.Add(1)
		go func(seq int) {
			defer wg.Done()
			for {
				select {
				case <-chs[seq]:
					fmt.Println(seq)
					time.Sleep(1 * time.Second)
					if seq == n-1 {
						chs[0] <- struct{}{}
					} else {
						chs[seq+1] <- struct{}{}
					}
				case <-exitCh:
					fmt.Printf("%d exit", seq)
					return
				}
			}
		}(i)
	}
	chs[0] <- struct{}{}
	<-sigCh
	close(exitCh)
	wg.Wait()
}
