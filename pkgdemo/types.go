package pkgdemo

import (
	"fmt"
	"unsafe"
)

func boolTest() {

}

func typeSize() {
	fmt.Println("map size:", unsafe.Sizeof(map[int]int{}))
	fmt.Println("channel size:", unsafe.Sizeof(make(chan int)))
	fmt.Println("slice size:", unsafe.Sizeof([]int{}))
}
