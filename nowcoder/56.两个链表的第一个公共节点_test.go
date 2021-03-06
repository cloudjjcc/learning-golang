package nowcoder

import (
	"fmt"
	"github.com/cloudjjcc/go-exercises/datastructures"
	"testing"
)

func Test_firstCommonNode(t *testing.T) {
	list1 := datastructures.BuildList([]int{1, 2, 3, 4})
	fmt.Println(firstCommonNode(list1, list1.Next).Value)
}
