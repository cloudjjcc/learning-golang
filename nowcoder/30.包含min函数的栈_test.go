package nowcoder

import (
	"fmt"
	"testing"
)

func TestNewMyStack(t *testing.T) {
	s := NewMyStack()
	s.Push(10)
	s.Push(2)
	s.Push(3)
	s.Push(3)
	s.Push(1)
	s.Push(9)
	fmt.Println(s.Min())
	s.Pop()
	s.Pop()
	fmt.Println(s.Min())
}
