package pkgdemo

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	arr := [0]int{}
	arr2 := [1]int{}
	fmt.Printf("%p,%p\n", &arr, &arr2)
	var i int
	fmt.Printf("%p\n", &i)
	//ptr := uintptr(0x1852ed8)
	//intp := (*[]int)(unsafe.Pointer(ptr))
	//fmt.Printf("%v,%v\n", intp, *intp)
	s := new(struct{})
	fmt.Printf("%p,%p\n", s, &struct{}{})
	var sli []int //nil 切片
	fmt.Printf("%v,%p,%p\n", sli, sli, &sli)
	sli2 := []int{} //空切片
	fmt.Printf("%v,%p,%p\n", sli2, sli2, &sli2)
	sli3 := make([]int, 0) // 空切片
	fmt.Printf("%v,%p,%p\n", sli3, sli3, &sli3)
	sli4 := new([]int) //nil 切片
	fmt.Printf("%v,%p,%p\n", sli4, *sli4, sli4)
	var pp *int
	fmt.Printf("%p\n", pp)
}
