package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//写代码实现两个 goroutine，
//其中一个产生随机数并写入到 go channel 中，
//另外一个从 channel 中读取数字并打印到标准输出。
//最终输出五个随机数。

func main() {
	numChan := make(chan int, 5)
	works := sync.WaitGroup{}
	works.Add(2)
	randomProducer := func() {
		defer works.Done()
		rand.Seed(time.Now().Unix())
		for i := 0; i < 5; i++ {
			numChan <- rand.Int()
		}
		close(numChan)
	}
	printNum := func() {
		defer works.Done()
		for num := range numChan {
			fmt.Println(num)
		}
	}
	go randomProducer()
	go printNum()
	works.Wait()
}
