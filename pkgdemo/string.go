package pkgdemo

import (
	"fmt"
	"unicode/utf8"
)

// string 相关知识点

// rune
// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
//type rune = int32

func runeLit() {
	ru := '哈'
	fmt.Printf("%T,%c\n", ru, ru)
}

func strLit() {
	strA := "hello"
	strB := `
		小明
`
	fmt.Print(strA, strB)
}

// string 的遍历
func strRange() {
	for i, v := range "hi,我是小明" {
		fmt.Printf("i:%d,v:%c,type v:%T\n", i, v, v)
	}
}

// string 的len
// string 没有cap方法
func strLen() {
	demoStr := "hi,我是小明"
	lenStr := len(demoStr)
	lenBytes := len([]byte(demoStr))
	countRunes := utf8.RuneCountInString(demoStr)
	lenRunes := len([]rune(demoStr))
	fmt.Printf("lenStr:%d,lenBytes:%d,countRunes:%d,lenRunes:%d\n", lenStr, lenBytes, countRunes, lenRunes)
}
