package nowcoder

import "fmt"

type anode struct {
	data int
	next *anode
}

func (a *anode) String() string {
	var str string
	for cur := a; cur != nil; cur = cur.next {
		str += fmt.Sprintf("->%d", cur.data)
	}
	return str
}

// 迭代法
func Merge(list1, list2 *anode) *anode {
	// check nil list
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// result list head
	nlist := new(anode)
	var (
		p1, p2, p3 = list1, list2, nlist
	)
	for p1 != nil && p2 != nil {
		if p1.data > p2.data {
			p3.next = p2
			p2 = p2.next
		} else {
			p3.next = p1
			p1 = p1.next
		}
		p3 = p3.next
	}
	if p1 == nil {
		p3.next = p2
	} else {
		p3.next = p1
	}
	return nlist.next
}
