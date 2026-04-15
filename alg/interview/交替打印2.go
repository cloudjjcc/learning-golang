package main

import (
	"strconv"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	evenChan := make(chan struct{})
	oddChan := make(chan struct{})
	i := 0
	next := func() string {
		if i == 100 {
			return ""
		}
		i++
		return strconv.Itoa(i)
	}
	printEven := func() {
		defer func() {
			close(oddChan)
			wg.Done()
		}()
		for range evenChan {
			s := next()
			if s == "" {
				return
			}
			println("even:" + s)
			oddChan <- struct{}{}
		}
	}
	printOdd := func() {
		defer func() {
			close(evenChan)
			wg.Done()
		}()
		for range oddChan {
			s := next()
			if s == "" {
				return
			}
			println("odd:" + s)
			evenChan <- struct{}{}
		}

	}
	go printEven()
	go printOdd()
	oddChan <- struct{}{}
	wg.Wait()
}
