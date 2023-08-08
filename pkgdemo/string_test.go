package pkgdemo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestStrAdd(t *testing.T) {
	str1 := "hello" + "world"
	str2 := "hello" + `world`
	c := 'a' + 'b'
	fmt.Printf("str1:%s\nstr2:%s\nc:%v", str1, str2, c)
}
func BenchmarkConcatWithAdd(b *testing.B) {
	strA := "hello"
	strB := "world"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = strA + strB
	}
}
func BenchmarkConcatWithSprintf(b *testing.B) {
	strA := "hello"
	strB := "world"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", strA, strB)
	}
}
func TestStrJoin(b *testing.T) {
	strA := "hello"
	strB := "world"
	strC := strings.Join([]string{strA, strB}, "")
	fmt.Println(strC)
}
func BenchmarkConcatWithBuffer(b *testing.B) {
	strA := "hello"
	buffer := bytes.Buffer{}
	buffer.Grow(b.N * 5)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buffer.WriteString(strA)
	}
}
func BenchmarkConcatWithBuilder(b *testing.B) {
	strA := "hello"
	builder := strings.Builder{}
	builder.Grow(b.N * 5)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		builder.WriteString(strA)
	}
}
func Test_runeLit(t *testing.T) {
	runeLit()
}

func Test_strLit(t *testing.T) {
	strLit()
}

func Test_strRange(t *testing.T) {
	strRange()
}

func Test_strLen(t *testing.T) {
	strLen()
}

func Test_typeConvention(t *testing.T) {
	str := "hello world"
	fmt.Printf("%T,%s", str[:], []byte(str))
}
func TestRuneToString(t *testing.T) {
	fmt.Println(string(rune(65)), string(rune(27494)))
}

func TestStrSlice(t *testing.T) {
	str := "hi,小明"
	str1 := str[:4]
	fmt.Printf("%T,%v", str1, str1)
}

func TestStrSize(t *testing.T) {
	str := "hi"
	str2 := "hello world"
	str3 := "小明"
	fmt.Printf("%d,%d,%d\n", unsafe.Sizeof(str), unsafe.Sizeof(str2), unsafe.Sizeof(str3))
	fmt.Printf("%p,%p,%p\n", &str, &str2, &str3)
}
func TestStrIndex(t *testing.T) {
	str := "hello world"
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&str))
	sh.Cap = len(str)
	s := *(*[]byte)(unsafe.Pointer(&sh))
	fmt.Printf("%s\n", s)
	str1 := "abc"
	str2 := "abc"
	fmt.Printf("%p,%p,%p,%p\n", unsafe.StringData(str1), unsafe.StringData(str2), &str1, &str2)
}

func TestRune(t *testing.T) {
	runes := []rune{'a', '好'}
	fmt.Printf("%s", string(runes))
}
