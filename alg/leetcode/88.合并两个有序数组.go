package leetcode

// 正向双指针 时间复杂度O(m+n) 空间复杂度O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := 0, 0
	tmp := make([]int, 0, m+n)
	for {
		if i1 >= m {
			tmp = append(tmp, nums2[i2:]...)
			break
		}
		if i2 >= n {
			tmp = append(tmp, nums1[i1:]...)
			break
		}
		if nums1[i1] < nums2[i2] {
			tmp = append(tmp, nums1[i1])
			i1++
		} else {
			tmp = append(tmp, nums2[i2])
			i2++
		}
	}
	copy(nums1, tmp)
}

// 逆向双指针O(m+n) O(1)
func merge1(nums1 []int, m int, nums2 []int, n int) {
	i1, i2 := m-1, n-1
	i := m + n - 1
	for {
		if i1 < 0 {
			copy(nums1, nums2[:i2+1])
			break
		}
		if i2 < 0 {
			break
		}
		if nums1[i1] > nums2[i2] {
			nums1[i] = nums1[i1]
			i1--
		} else {
			nums1[i] = nums2[i2]
			i2--
		}
		i--
	}
}
