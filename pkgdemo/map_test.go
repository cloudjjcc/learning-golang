package pkgdemo

import (
	"fmt"
	"testing"
)

func TestMapInit(t *testing.T) {
	m := map[string]int{}
	var m1 map[string]int
	m2 := make(map[string]int)
	fmt.Printf("m:%p,m1:%p,m2:%p\n", m, m1, m2)
	m["math"] = 90
	m1["math"] = 90 //panic: assignment to entry in nil map
	m2["math"] = 90
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
		fmt.Printf("math score is:%d", v)
	}

}
