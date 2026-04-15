package main

import "fmt"

func main() {
	var a string
	//_, err := fmt.Scanln(&a)
	//if err != nil {
	//	return
	//}
	//fmt.Println(a)
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}
	fmt.Println(a)
	//var b int
	//_, err := fmt.Scanf("%d\n", &b)
	//if err != nil {
	//	return
	//}
	//fmt.Println(b)
}
