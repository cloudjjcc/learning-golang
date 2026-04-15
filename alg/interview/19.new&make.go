package main

import "fmt"

func main() {
	testNew()
}

func testNew() {
	i := new(int)
	fmt.Printf("address is %p,value is %v\n", i, *i)
	s := new(struct{})
	fmt.Printf("address is %p,value is %v\n", s, *s)

}
