package pkgdemo

import (
	"fmt"
	"unsafe"
)

func boolTest() {

}

func typeSize() {
	// 不好的布局：有内存浪费
	type BadLayout struct {
		a bool  // 1 byte
		b int64 // 8 bytes (需要对齐，浪费7字节)
		c bool  // 1 byte (浪费7字节)
	} // 总大小: 24 bytes

	// 好的布局：紧凑排列
	type GoodLayout struct {
		b int64 // 8 bytes
		a bool  // 1 byte
		c bool  // 1 byte
		// 6 bytes padding (只浪费6字节)
	} // 总大小: 16 bytes

	fmt.Printf("BadLayout size:  %d\n", unsafe.Sizeof(BadLayout{}))  // 24
	fmt.Printf("GoodLayout size: %d\n", unsafe.Sizeof(GoodLayout{})) // 16

	// 查看字段偏移
	var good GoodLayout
	fmt.Printf("b offset: %d\n", unsafe.Offsetof(good.b)) // 0
	fmt.Printf("a offset: %d\n", unsafe.Offsetof(good.a)) // 8
	fmt.Printf("c offset: %d\n", unsafe.Offsetof(good.c)) // 9
	fmt.Println("map size:", unsafe.Sizeof(map[int]int{}))
	fmt.Println("channel size:", unsafe.Sizeof(make(chan int)))
	fmt.Println("slice size:", unsafe.Sizeof([]int{}))
	fmt.Println("interface size:", unsafe.Sizeof(any(nil)))
	fmt.Println("func size:", unsafe.Sizeof(func() {}))
	fmt.Printf("func addr:%p\n", func() {})
}
