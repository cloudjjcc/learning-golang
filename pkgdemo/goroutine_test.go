package pkgdemo

import (
	"fmt"
	"testing"
)

func TestCreateNew(t *testing.T) {
	//go (func())(nil)()//fatal error: go of nil func value
	var ints [512]int
	//go func([512]int) {}(ints)//fatal error: newproc: function arguments too large for new goroutine
	go func() {
		fmt.Println(ints)
	}()
}
