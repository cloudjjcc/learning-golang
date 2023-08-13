package pkgdemo

import (
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("bey bey")
	fmt.Println("do something ...")
}

func TestDefer2(t *testing.T) {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	time.Sleep(1 * time.Second)
}

func TestDefer3(t *testing.T) {
	defer ((func())(nil))()
	fmt.Println("do something ...")
}

type tt struct {
	sum int
}

func (t *tt) add(a int) int {
	t.sum += a
	return t.sum
}
func TestMethod(t *testing.T) {
	ts := tt{}
	add := (*tt).add
	add(&ts, 1)
	add2 := ts.add
	add2(2)
	fmt.Printf("%T,%T,%v\n", add, add2, ts.sum)
}
