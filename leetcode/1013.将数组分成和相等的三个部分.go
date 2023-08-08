package leetcode

//给你一个整数数组 A，只有可以将其划分为三个和相等的非空部分时才返回 true，否则返回 false。
//形式上，如果可以找出索引 i+1 < j 且满足 (A[0] + A[1] + ... + A[i] == A[i+1] + A[i+2] + ... + A[j-1] == A[j] + A[j-1] + ... + A[A.length - 1]) 就可以将数组三等分。
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func canTreePartsEqualSum(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	if sum%3 != 0 {
		return false
	}
	psum := sum / 3
	tsum := 0
	left := 0
	for i := 0; i < len(arr); i++ {
		tsum += arr[i]
		if tsum == psum {
			left = i
			break
		}
	}
	tsum = 0
	for i := left + 1; i < len(arr)-1; i++ {
		tsum += arr[i]
		if tsum == psum {
			return true
		}
	}
	return false
}
