package pkgdemo

import (
	"fmt"
	"testing"
)

func TestStructAddr(t *testing.T) {
	stu := Student{}
	fmt.Printf("%p,%+v\n", &stu, stu)
	var stu1 *Student
	fmt.Printf("%p,%p,%+v\n", &stu1, stu1, stu1)
	stu2 := &Student{}
	fmt.Printf("%p,%p,%+v\n", &stu2, stu2, stu2)
}

func TestStructMethod(t *testing.T) {
	stu := &Student{Name: "xiaoming"}
	fmt.Println(stu.GetName())
	fmt.Println((*Student).GetName(stu))
}
