package pkgdemo

import (
	"fmt"
	"testing"
)

func TestArr(t *testing.T) {
	arr := [...]int{1, 2, 3}
	fmt.Printf("len:%d,cap:%d", len(arr), cap(arr))
}
func TestMakeSlice(t *testing.T) {
	arr := [...]int{1, 2, 3}
	sli := arr[:]
	sli2 := sli[1:]
	fmt.Printf("arr:(%p,%T),sli:(%p,%T),sli2:(%p,%T)", &arr, arr, sli, sli, sli2, sli2)
}

func TestAppend(t *testing.T) {
	sli := []int{1, 2, 3}
	fmt.Printf("sli(data_addr:%p，type_addr:%p,cap:%d)\n", sli, &sli, cap(sli))
	sli = append(sli, 4)
	fmt.Printf("sli(data_addr:%p，type_addr:%p,cap:%d)\n", sli, &sli, cap(sli))
	sli2 := append(sli, 5)
	fmt.Printf("sli2:(data_addr:%p，type_addr:%p,cap:%d)\n", sli2, &sli2, cap(sli2))
}

func TestCopy(t *testing.T) {
	src := []string{"hello", "world", "小明"}
	dst := make([]string, 3)
	copy(dst, src)
	fmt.Printf("src_data_addr:%p,dst_data_addr:%p\n", src, dst)
	fmt.Println(dst)
}
