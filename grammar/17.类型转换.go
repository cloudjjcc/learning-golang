package main

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"reflect"
	"unsafe"
)

type animal interface {
	Foot() int
}
type Bird struct {
}

func (b Bird) Foot() int {
	return 2
}

func main() {
	var (
		ani1 animal = Bird{}
		ani2 animal = &Bird{}
		ani3 animal = new(Bird)
		stu  struct{}
		inf  interface{}
	)
	ani1.Foot()
	ani2.Foot()
	ani3.Foot()
	fmt.Println(reflect.TypeOf(ani1), unsafe.Sizeof(ani1))
	fmt.Println(reflect.TypeOf(ani2), unsafe.Sizeof(ani2))
	fmt.Println(reflect.TypeOf(ani3), unsafe.Sizeof(ani3))
	fmt.Println(reflect.TypeOf(stu), unsafe.Sizeof(stu))
	fmt.Println(reflect.TypeOf(inf), unsafe.Sizeof(inf))
	// nil
	var nilObj *datastructures.SingleList
	fmt.Println(nilObj)
	fmt.Println(nilObj.Len())
}
