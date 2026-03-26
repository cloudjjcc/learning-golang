package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	addressM := make(map[*Node]*Node)
	// 复制val
	for cur := head; cur != nil; cur = cur.Next {
		t := new(Node)
		t.Val = cur.Val
		addressM[cur] = t
	}
	// 复制Next&&Random
	for src, dst := range addressM {
		dst.Next = addressM[src.Next]
		dst.Random = addressM[src.Random]
	}
	return addressM[head]
}
