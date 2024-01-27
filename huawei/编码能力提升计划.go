package huawei

func mainminTime() {

}

func minTime(time []int, m int) int {
	left, right, mid := 0, 0, 0
	for _, v := range time {
		right += v
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	check := func(v int) bool {
		usedDay, totalTime, maxTime := 1, 0, 0
		for i := 0; i < len(time); i++ {
			nextTime := min(maxTime, time[i])
			if nextTime+totalTime <= v {
				totalTime += nextTime
				maxTime = max(maxTime, time[i])
			} else {
				usedDay++
				totalTime = 0
				maxTime = time[i]
			}
		}
		return usedDay <= m
	}
	for left <= right {
		mid = (left + right) >> 1
		if check(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
