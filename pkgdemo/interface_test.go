package pkgdemo

import (
	"fmt"
	"testing"
)

func TestInterfaceNil(t *testing.T) {
	var a interface{}
	fmt.Println(a == nil)
	var b *int
	a = b
	fmt.Println(a == nil)
	var m map[string]int
	i := funcnil(m)
	fmt.Println(i == nil)
}

func funcnil(a interface{}) interface{} {
	fmt.Println(a == nil)
	m := a.(map[string]int)
	fmt.Println(m == nil)
	return a
}
