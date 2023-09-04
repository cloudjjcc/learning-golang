package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
	age  int
}

var p Person

//var mu sync.Mutex
var (
	updateCh = make(chan Person)
	closeCh  = make(chan struct{})
)

func update(name string, age int) {
	fmt.Printf("update start %s,%d\n", name, age)
	p.name = name
	time.Sleep(200 * time.Millisecond)
	p.age = age
	fmt.Printf("update finish %s,%d\n", name, age)
}
func main() {

	wg := sync.WaitGroup{}
	wg.Add(11)
	go func() {
		defer wg.Done()
		for ch := range updateCh {
			update(ch.name, ch.age)
		}
	}()
	for i := 0; i < 10; i++ {
		name, age := fmt.Sprintf("name:%d", i), i
		update(name, age)
	}
	wg.Wait()
	fmt.Printf("name:%s,age:%d\n", p.name, p.age)
}
