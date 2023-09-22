package pkgdemo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMapInit(t *testing.T) {
	//m := map[string]int{}
	var m1 map[string]int
	fmt.Printf("%p,len:%d,size:%v,is nil:%t\n", m1, len(m1), unsafe.Sizeof(m1), m1 == nil)
	fmt.Println(m1["a"])
	//m1["a"] = 2
	m2 := map[int]string{
		1: "cat",
		2: "dog",
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	//m2 := make(map[string]int)
	//fmt.Printf("m:%p,m1:%p,m2:%p\n", m, m1, m2)
	//m["math"] = 90
	//m1["math"] = 90 //panic: assignment to entry in nil map
	//m2["math"] = 90
	m := make(map[string]string)
	m["a"] = "abcddd"
	m3 := make(map[string]string, 100)
	m3["a"] = "dsdsdds"
	m3["b"] = "dadsd"
	fmt.Printf("%d,%d", len(m), len(m3))
}

func TestOp(t *testing.T) {
	m := map[string]int{
		"english": 80,
	}
	m["math"] = 90
	m["math"]++
	//addr:=&m["math"]
	fmt.Println(len(m))
	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}
	delete(m, "math")
	if v, ok := m["math"]; ok {
		fmt.Printf("math scoe is:%d", v)
	}

}
func TestMapAddr(t *testing.T) {
	m := make(map[int]int)
	fn := func(mm map[int]int) {
		fmt.Printf("%p,%p\n", mm, &mm)
	}
	fn(m)
	fmt.Printf("%p,%p\n", m, &m)
}

func TestMapHeader(t *testing.T) {
	//m := make(map[int]byte)
	//m[1] = 'a'
	//m[2] = 'b'
	//m[3] = 'c'
	//m[4] = 'd'
	//h := *(**hmap)(unsafe.Pointer(&m))
	//v1 := h.buckets
	//v2 := (*bmap)(unsafe.Add(unsafe.Pointer(h.buckets), unsafe.Sizeof(*v1)))
	//v3 := (*bmap)(unsafe.Add(unsafe.Pointer(h.buckets), 2*unsafe.Sizeof(*v1)))
	//fmt.Printf("%v,%v,%v,size:%v\n", v1, v2, v3, unsafe.Sizeof(m))

}

// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    *bmap          // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra uintptr // optional fields
}

// A bucket for a Go map.
type bmap struct {
	topbits [8]uint8
	keys    [8]int
	values  [8]byte
	pad     uintptr
	//overflow uintptr
}
