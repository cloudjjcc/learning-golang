package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//问题描述
//使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func main() {
	var (
		numCh           = make(chan struct{}, 1)
		letterCh        = make(chan struct{}, 1)
		i         int32 = 0
		waitGroup sync.WaitGroup
	)
	printNum := func() {
		defer waitGroup.Done()
		for {
			select {
			case <-numCh:
				ii := atomic.LoadInt32(&i)
				fmt.Printf("%d%d", ii*2+1, ii*2+2)
				letterCh <- struct{}{}
				if ii == 13 {
					return
				}
			default:
			}
		}
	}
	printLetter := func() {
		defer waitGroup.Done()
		for {
			select {
			case <-letterCh:
				ii := atomic.LoadInt32(&i)
				if ii == 13 {
					return
				}
				fmt.Printf("%c%c", ii*2+1+64, ii*2+2+64)
				atomic.AddInt32(&i, 1)
				numCh <- struct{}{}
			default:
			}

		}
	}
	numCh <- struct{}{}
	waitGroup.Add(2)
	go printNum()
	go printLetter()
	waitGroup.Wait()
}
