package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//testWaitGroup1()
	//testWaitGroup2()
	//testMutex1()
}

//assignment copies lock value to yawg: sync.Mutex
//call of fmt.Println copies lock value: sync.Mutex
//call of fmt.Println copies lock value: sync.Mutex
func testMutex1() {
	//var wg sync.Mutex
	//yawg:=wg
	//fmt.Println(wg, yawg)
}

func testWaitGroup2() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("goroutine run ....")
	}(&wg)
	wg.Wait()
}

func testWaitGroup1() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: -------", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:++++++++ ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
