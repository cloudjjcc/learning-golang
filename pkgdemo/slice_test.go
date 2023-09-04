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
func TestSliceCap(t *testing.T) {
	slice := make([]int, 0, 8)
	sli := slice[2:3]
	sli2 := sli[1:5]
	fmt.Println(sli2)
}

func TestSliceOp(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:0]
	t.Logf("%v,%v,%p,%p", s1, s2, s1, s2)
	t.Log(cap(s2))
	s3 := s2[0:0:2]
	t.Log(cap(s3))
}

func TestSliceAppend(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}
	t.Logf("%p,len:%d,cap:%d", s1, len(s1), cap(s1))
	_ = append(s1, 7)
	t.Logf("%p,len:%d,cap:%d", s1, len(s1), cap(s1))
	s1 = append(s1, 7)
	t.Logf("%p,len:%d,cap:%d", s1, len(s1), cap(s1))
	s2 := s1[1:2]
	t.Logf("%p,len:%d,cap:%d", s2, len(s2), cap(s2))
}

func TestSliceEqual(t *testing.T) {
	//s1:=[]int{1,2,3}
	//s2:=[]int{1,2,3}
	//t.Log(s1==s2)
}
