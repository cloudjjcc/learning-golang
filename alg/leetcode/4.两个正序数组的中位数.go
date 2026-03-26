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

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	k := (m + n) / 2
	i, j := 0, 0
	t := 0
	t1, t2 := 0, 0
	for {
		if i >= m {
			if t == k-1 {
				t1 = nums2[j]
			}
			if t == k {
				t2 = nums2[j]
				break
			}
			j++
			t++

			continue
		}
		if j >= n {
			if t == k-1 {
				t1 = nums1[i]
			}
			if t == k {
				t2 = nums1[i]
				break
			}
			i++
			t++
			continue
		}
		if nums1[i] < nums2[j] {
			if t == k-1 {
				t1 = nums1[i]
			}
			if t == k {
				t2 = nums1[i]
				break
			}
			i++
			t++

		} else {
			if t == k-1 {
				t1 = nums2[j]
			}
			if t == k {
				t2 = nums2[j]
				break
			}
			j++
			t++
		}
	}
	if (m+n)%2 == 0 {
		return float64(t1+t2) / float64(2)
	}
	return float64(t2)
}
