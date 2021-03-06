package nowcoder

import (
	"fmt"
	"testing"
)

func Test_push(t *testing.T) {
	push(1)
	push(2)
	push(3)
	pop()
	pop()
	push(4)
	fmt.Println(pop())
}
