package nowcoder

import (
	"fmt"
	"testing"
)

func Test_serializable(t *testing.T) {
	testTreeStr := "0!#1!#2!#3!#4!#"
	tree := deserializable([]byte(testTreeStr))
	fmt.Println(string(serializable(tree)))
}
