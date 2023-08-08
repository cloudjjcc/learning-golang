package leetcode

// 利用map O(m+n) O(m+n)
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, v := range nums1 {
		m1[v] = 1
	}
	var ans []int
	for _, v := range nums2 {
		if m1[v] > 0 {
			ans = append(ans, v)
			m1[v] = 0
		}
	}
	return ans
}

// 排序加双指针O(mlogm+nlogn) O(logm+logn) TODO
func intersection1(nums1 []int, nums2 []int) []int {
	return nil
}
