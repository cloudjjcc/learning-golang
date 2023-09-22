package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	// 尝试从任意加油站出发
	for i := 0; i < len(gas); i++ {
		n := 0
		left := 0
		for {
			// 加油
			left += gas[(i+n)%len(gas)]
			left -= cost[(i+n)%len(gas)]
			if left < 0 {
				break
			}
			n++
			if n == len(gas) {
				return i
			}
		}
	}
	return -1
}
