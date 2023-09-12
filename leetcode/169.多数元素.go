package leetcode

func majorityElement(nums []int) int {
	condidate := nums[0]
	cnt := 0
	for _, v := range nums {
		if cnt == 0 {
			condidate = v
		}
		if v == condidate {
			cnt++
			continue
		}
		cnt--
	}
	return condidate
}
