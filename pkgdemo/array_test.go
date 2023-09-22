package pkgdemo

import (
	"fmt"
	"testing"
)

func TestArrayType(t *testing.T) {
	var (
		a [2]int
		b [3]int
		c [2]interface{}
		d [3]interface{}
	)
	fmt.Printf("a type:%T,b type:%T,c type:%T,d type:%T\n", a, b, c, d)
	e := [...]int{1, 2, 3}
	fmt.Printf("%T\n", e)
	f := [8]int{}
	fmt.Printf("len:%d,cap:%d\n", len(f), cap(f))
}
