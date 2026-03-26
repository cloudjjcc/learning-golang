package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// test compare
	//compare()
	//testAppend()
	//sliceSize()
	//arrSize()
	//sliceOP()
	//arr:=make([]int,10)
	//arr:=[]string{}
	//fmt.Println(arr)
	strs := []string{}

	fmt.Println(strs[0])
	fmt.Println(unsafe.Sizeof(strs))
}

func sliceOP() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7}
	sli := arr[1:2:3]
	fmt.Printf("sli is %v,len is %d,cap is %d\n", sli, len(sli), cap(sli))
}

func arrSize() {
	var tmp = [4]int{1, 2, 3, 4}
	fmt.Println(unsafe.Sizeof(tmp))
}

func sliceSize() {
	var tmp []int
	fmt.Println(unsafe.Sizeof(tmp))
}
func compare() {
	str1 := []string{"a", "b", "c"}
	str2 := str1[1:]
	str2[1] = "new"
	fmt.Println(str1)
	fmt.Printf("str1:%p,str2:%p\n", str1, str2)
	str2 = append(str2, "z", "x", "y")
	fmt.Println(str1)
	fmt.Printf("str1:%p,str2:%p\n", str1, str2)
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("nilSlice:%v,cap:%d;emptySlice:%v,cap:%d\n", nilSlice, cap(nilSlice), emptySlice, cap(emptySlice))
	fmt.Println([...]string{"1"} == [...]string{"1"})
	// 切片不能直接比较
	//fmt.Println([]string{"1"}==[]string{"1"})
}

func testAppend() {
	buf := make([]int, 1, 10)
	for i := 1; i < 1000; i++ {
		buf = append(buf, i)
		fmt.Printf("len %d,address %p\n", len(buf), buf)
	}
}
