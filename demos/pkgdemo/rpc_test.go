package pkgdemo

import "testing"

func Test_serve(t *testing.T) {
	go serve()
	client()
}
