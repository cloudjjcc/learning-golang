package nowcoder

//题目描述
//输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。
//如果是则输出Yes,否则输出No。
//假设输入的数组的任意两个数字都互不相同。

func verifySequenceOfBST(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	rootVal := arr[len(arr)-1]
	// left node < root node
	i := 0
	for i = 0; i < len(arr)-1; i++ {
		if arr[i] > rootVal {
			break
		}
	}
	// right node >= root node
	for j := i; j < len(arr)-1; j++ {
		if arr[j] < rootVal {
			return false
		}
	}
	// verify left child tree and right child tree
	return verifySequenceOfBST(arr[:i]) && verifySequenceOfBST(arr[i:len(arr)-1])
}
