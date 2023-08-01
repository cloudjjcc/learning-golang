package leetcode

import "math"

// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
//
// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
//
// 任何误差小于 10-5 的答案都将被视为正确答案。
//
// 示例 1：
//
// 输入：nums = [1,12,-5,-6,50,3], k = 4
// 输出：12.75
// 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
// 示例 2：
//
// 输入：nums = [5], k = 1
// 输出：5.00000
//
// 提示：
//
// n == nums.length
// 1 <= k <= n <= 105
// -104 <= nums[i] <= 104
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/maximum-average-subarray-i
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	sum := 0
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if i >= k {
			sum -= nums[i-k]
			if sum > max {
				max = sum
			}
		} else {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
