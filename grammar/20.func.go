package main

import (
	"errors"
	"fmt"
	"reflect"
)

type T interface {
	a(a int, b string) interface{}
}

func main() {
	var fn func(a int, b string) interface{}
	errors.Join()
	fmt.Println(reflect.TypeOf(fn).Implements(reflect.TypeOf((T)(nil))))

}
