package main

import (
	"fmt"
	"sync"
)

func main() {
	lock1 := &sync.Mutex{}
	lock2 := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			lock1.Lock()
			defer lock1.Unlock()
			lock2.Lock()
			defer lock2.Unlock()
			fmt.Println("a running")
		}

	}()
	go func() {
		defer wg.Done()
		for {
			lock2.Lock()
			defer lock2.Unlock()
			lock1.Lock()
			defer lock1.Unlock()
			fmt.Println("b running")
		}

	}()
	wg.Wait()
}
