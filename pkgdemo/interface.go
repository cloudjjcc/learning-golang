package pkgdemo

var _ error = (*MyError)(nil)

type MyError struct {
	error
}
