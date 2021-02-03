package pkgdemo

import (
	"archive/tar"
	"bytes"
	"fmt"
)

func demoZip() {
}

func demoTar() {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	tw.Write([]byte("hello world"))
	header := new(tar.Header)
	tw.WriteHeader(header)
	tr := tar.NewReader(buf)
	fmt.Println(tr.Next())
}
