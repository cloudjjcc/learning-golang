package pkgdemo

import (
	"fmt"
)

func ArrayTest() {
	arr := [4]int{}
	fmt.Printf("%T，%T\n", arr, &arr)
	fmt.Println(len(arr))
	fmt.Printf("%p,%p,%p\n", &arr, &arr[0], &arr[1])
}
