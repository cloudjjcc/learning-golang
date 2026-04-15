package nowcoder

import (
	"fmt"
	ds "github.com/cloudjjcc/learning-golang/alg/datastructures"
	"testing"
)

func Test_convert(t *testing.T) {
	tree := ds.BuildTreeFromArray([]interface{}{10, 8, 12, 7, 9, 11, 13})
	dlinkedlist := convert(tree)
	fmt.Println(dlinkedlist.Value)
}
