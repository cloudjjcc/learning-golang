package main

import (
	"fmt"
	"reflect"
)

type T interface {
	a(a int, b string) interface{}
}

func main() {
	var fn func(a int, b string) interface{}
	fmt.Println(reflect.TypeOf(fn).Implements(reflect.TypeOf((T)(nil))))

}
