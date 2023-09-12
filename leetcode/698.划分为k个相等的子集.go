package leetcode

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	//tsum := sum / k
	// TODO
	return false
}
