package pkgdemo

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_boolTest(t *testing.T) {
	fmt.Printf("%d", unsafe.Sizeof(true))
}
