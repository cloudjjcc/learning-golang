package pkgdemo

import (
	"fmt"
)

func updateValue() {
	var val int
	var val1, val2 interface{}
	val1 = val
	val2 = val1
	fmt.Println(val1, val2)

}
