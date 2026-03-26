package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	catCh, dogCh, fishCh := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-catCh
			fmt.Println("cat:" + strconv.Itoa(i))
			dogCh <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-dogCh
			fmt.Println("dog:" + strconv.Itoa(i))
			fishCh <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			<-fishCh
			fmt.Println("fish:" + strconv.Itoa(i))
			catCh <- struct{}{}
		}
	}()
	catCh <- struct{}{}
	wg.Wait()
}
