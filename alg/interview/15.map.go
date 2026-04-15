package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
	"runtime"
	"time"
	"unsafe"
)

type Student struct {
	Age int
}

func main() {

	//benchmarkMap()
	//mapSize()
	//mapKey()
	//mapAccess1()
	newMap()
}

type mk struct {
	Value string
}

func mapKey() {
	var tmp = map[mk]int{
		mk{Value: "1"}: 1,
	}
	fmt.Println(tmp[mk{Value: "1"}])

}

func mapSize() {
	var tmp map[string]int
	tmp = make(map[string]int)
	fmt.Println(unsafe.Sizeof(tmp))
}
func mapAccess1() {
	students := map[string]Student{"xiaoming": {19}}
	// 如果map的value是结构体则无法直接对结构体的属性进行赋值
	//students["xiaoming"].Age = 18
	if stu, ok := students["xiaoming"]; ok {
		stu.Age = 20
		fmt.Println(stu)
	}
}

type wrapInt struct {
	value int
}

func benchmarkMap() {
	//testMap:=make(map[int]*wrapInt)
	testHashMap := treemap.NewWith(func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	// insert
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println("mem:", mem.Alloc)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		testHashMap.Put(i, &wrapInt{i})
	}
	runtime.ReadMemStats(&mem)
	fmt.Println("insert cost :", time.Since(start), ",mem:", mem.Alloc)

	// get
	start = time.Now()
	for i := 0; i < 10000000; i++ {
		if _, ok := testHashMap.Get(i); ok {

		}
	}
	fmt.Println("get cost :", time.Since(start))

	fmt.Println("map len:", testHashMap.Size())
}

func newMap() {
	// map 不可以用new分配
	tmp := new(map[int]int)
	(*tmp)[1] = 100
}
