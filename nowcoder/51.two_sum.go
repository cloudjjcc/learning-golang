package nowcoder

//题目描述
//给出一个整数数组，请在数组中找出两个加起来等于目标值的数，
//你给出的函数twoSum 需要返回这两个数字的下标（index1，index2），需要满足 index1 小于index2.。注意：下标是从1开始的
//假设给出的数组中只存在唯一解
//例如：
//给出的数组为 {2, 7, 11, 15},目t标值为9
//输出 ndex1=1, index2=2

// 双重遍历 O(n*n)
func twoSum(arr []int, target int) []int {
	for i, v := range arr {
		for j := i; j < len(arr); j++ {
			if v == target-arr[j] {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{-1, -1}
}

// 借助map O(n)
func twoSum2(arr []int, target int) []int {
	m := make(map[int]int)
	for i, v := range arr {
		if j, ok := m[v]; ok {
			return []int{j + 1, i + 1}
		}
		m[target-v] = i
	}
	return []int{-1, -1}
}

// 双指针
