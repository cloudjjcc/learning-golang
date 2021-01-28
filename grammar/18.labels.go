package main

import "fmt"

func main() {
	// break 跳出标签循环
	testBreakLabel()
	// continue 跳转到标签循环继续执行剩下的循环
	//testContinueLabel()
	// infinite loop,goto 跳转到标签继续执行标签下的语句
	//testGotoLabel()
}

func testBreakLabel() {

	for i := 0; i < 10; i++ {
	loop2:
		for j := 0; j < 10; j++ {
			if j == 5 {
				break loop2
			}
			fmt.Println(i, j)
		}
	}
}
func testContinueLabel() {
loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				continue loop
			}
			fmt.Println(i, j)
		}
	}
}
func testGotoLabel() {
loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				goto loop
			}
			fmt.Println(i, j)
		}
	}
}
