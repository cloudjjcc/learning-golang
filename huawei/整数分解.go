package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	for {
		i := 0
		n, err := fmt.Scanln(&i)
		if n == 0 || err != nil {
			return
		}
		fmt.Println(getResult(i))
	}
}

func getResult(n int) string {
	buf := bytes.NewBufferString("")
	tmp := n
	for i := 2; i <= tmp; i++ {
		for tmp%i == 0 {
			buf.WriteString(strconv.Itoa(i))
			if tmp == i {
				buf.WriteString("=")
				buf.WriteString(strconv.Itoa(n))
			} else {
				buf.WriteString("*")
			}
			tmp /= i
		}
	}
	return buf.String()
}
