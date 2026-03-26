package nowcoder

type lnode struct {
	data interface{}
	next *lnode
}

func ReverseList(list *lnode) *lnode {
	if list == nil {
		return nil
	}
	var (
		prev *lnode
		cur  = list
		next *lnode
	)
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}
