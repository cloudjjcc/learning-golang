package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := md5.New()
	hash.Write([]byte("test"))
	sumBytes := hash.Sum([]byte{})
	fmt.Println(hex.EncodeToString(sumBytes))
}
