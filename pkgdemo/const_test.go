package pkgdemo

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	const (
		a, b = "golang", 100
		d, e
		f bool = true
		g
	)
	fmt.Println(d, e, g)
}
