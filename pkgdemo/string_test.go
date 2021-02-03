package pkgdemo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
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
