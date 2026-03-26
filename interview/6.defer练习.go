package main

import (
	"fmt"
	"time"
)

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
func main() {
	defer fmt.Println("finished")
	//cost()
	//defer_call()
	fmt.Println("start")
	//fval()
	pnTest()
	fmt.Println("end")
}

func cost() {
	defer func(t time.Time) { fmt.Println(time.Since(t)) }(time.Now())
	time.Sleep(1 * time.Second)
}

type tt struct {
	val string
}

func (s *tt) Say() {
	fmt.Println(s.val)
}
func fval() {
	t := tt{val: "aa"}
	defer t.Say()
	t = tt{val: "bb"}
}

func pnTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("test panic")
}
