package leetcode

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(r *TreeNode, pathSum int)
	path := make([]int, 0)
	ans := make([][]int, 0)
	dfs = func(r *TreeNode, pathSum int) {
		if r == nil {
			return
		}
		pathSum += r.Val
		path = append(path, r.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if r.Left == nil && r.Right == nil {
			if pathSum == targetSum {
				t := make([]int, len(path))
				copy(t, path)
				ans = append(ans, t)
			}
			return
		}
		if r.Left != nil {
			dfs(r.Left, pathSum)
		}
		if r.Right != nil {
			dfs(r.Right, pathSum)
		}
	}
	dfs(root, 0)
	return ans
}
