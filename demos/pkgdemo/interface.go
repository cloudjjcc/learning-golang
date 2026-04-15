package pkgdemo

import "fmt"

var _ error = (*MyError)(nil)
var _ MyInterface = (*MyError)(nil)

type MyError struct {
	error
	fmt.Stringer
}

type MyInterface interface {
	error
	fmt.Stringer
}
