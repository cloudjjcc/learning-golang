package pkgdemo

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_boolTest(t *testing.T) {
	fmt.Printf("%d", unsafe.Sizeof(true))
}

func TestStack(t *testing.T) {
	printK(0)
}
func printK(k int) {
	fmt.Println(k)
	printK(k + 1)
}
