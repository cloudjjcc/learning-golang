package leetcode

// 给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
//
// 示例 1：
//
// 输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
// 输出：6
// 解释：[1,1,1,0,0,1,1,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 6。
// 示例 2：
//
// 输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
// 输出：10
// 解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// 粗体数字从 0 翻转到 1，最长的子数组长度为 10。
//
// 提示：
//
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1
// 0 <= k <= nums.length
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode.cn/problems/max-consecutive-ones-iii
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	ans := 0
	zeros := 0 //区间内零的个数
	maxFn := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for right < len(nums) {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		ans = maxFn(ans, right-left+1)
		right++
	}
	return ans
}
