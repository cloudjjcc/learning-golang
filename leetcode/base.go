package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(arr []int) *ListNode {
	list := new(ListNode)
	if len(arr) == 0 {
		return list
	}
	list.Val = arr[0]
	if len(arr) == 1 {
		return list
	}
	cur := list
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return list
}

func ListToSlice(list *ListNode) []int {
	var ans []int
	for list != nil {
		ans = append(ans, list.Val)
		list = list.Next
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
