package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 合并数组
	ll := len(nums1) + len(nums2)
	if ll == 0 {
		return 0
	}
	nums3 := make([]int, 0, ll)
	a, b := 0, 0
	for i := 0; i < ll; i++ {
		if a >= len(nums1) { //nums1没有元素
			nums3 = append(nums3, nums2[b:]...)
			break
		}
		if b >= len(nums2) { //nums2没有元素
			nums3 = append(nums3, nums1[a:]...)
			break
		}
		if nums1[a] < nums2[b] {
			nums3 = append(nums3, nums1[a])
			a++
		} else {
			nums3 = append(nums3, nums2[b])
			b++
		}
	}
	mid := float64(0)
	if ll%2 == 0 {
		mid = float64(nums3[ll/2]+nums3[ll/2-1]) / 2
	} else {
		mid = float64(nums3[ll/2])
	}
	return mid
}
