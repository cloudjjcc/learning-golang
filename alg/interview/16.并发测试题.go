package main

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Mutex、RWMutex、WaitGroup、Once、Pool、Map
var mu sync.Mutex
var chain string
var rwmu sync.RWMutex
var count int

func main() {
	//fatal error: all goroutines are asleep - deadlock!
	//testMutex()
	//fatal error: all goroutines are asleep - deadlock!
	//testRWMutex()
	//panic: sync: WaitGroup is reused before previous Wait has returned
	//testWaitGroup()
	//testMyOnce()
	//testMyMutex()
	// 内存暴涨
	//testPool()
	// 一段时间后输出2
	//testChannel()
	//  close of nil channel
	//testChannel2()
	//testMap()
}

func testMap() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(m.Len())
}

func testChannel2() {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

func testChannel() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)
	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func testPool() {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}
	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}

type MyMutex struct {
	sync.Mutex
	count int
}

func testMyMutex() {
	var mu MyMutex
	mu.Lock()
	// 赋值之后的mu2为上锁状态
	var mu2 = mu
	mu.count++
	mu.Unlock()
	// 需要先解锁
	mu2.Unlock()
	mu2.Lock()
	mu2.count++
	mu2.Unlock()
	fmt.Println(mu.count, mu2.count)
}

type MyOnce struct {
	m    sync.Mutex
	done uint32
}

func (o *MyOnce) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}
func testMyOnce() {
	var once MyOnce
	testFunc := func() {
		fmt.Println("hello")
	}
	for i := 0; i < 100000; i++ {
		go once.Do(testFunc)
	}
}

func testWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		wg.Add(1)
	}()
	wg.Wait()
}
func testMutex() {
	chain = "main"
	A()
	fmt.Println(chain)
}
func A() {
	mu.Lock()
	defer mu.Lock()
	chain = chain + " --> A"
	B()
}
func B() {
	chain = chain + " --> B"
	C()
}
func C() {
	mu.Lock()
	defer mu.Lock()
	chain = chain + " --> C"
}
func testRWMutex() {
	go AA()
	time.Sleep(2 * time.Second)
	rwmu.Lock()
	defer rwmu.Unlock()
	count++
	fmt.Println(count)
}
func AA() {
	rwmu.RLock()
	defer rwmu.RUnlock()
	BB()
}
func BB() {
	time.Sleep(5 * time.Second)
	CC()
}
func CC() {
	rwmu.RLock()
	defer rwmu.RUnlock()
}
