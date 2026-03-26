package leetcode

func predictPartyVictory(senate string) string {
	var (
		rq []int
		dq []int
	)
	push := func(q *[]int, state int) {
		*q = append(*q, state)
	}
	pop := func(q *[]int) int {
		tmp := (*q)[0]
		*q = (*q)[1:]
		return tmp
	}
	for i, v := range senate {
		if v == 'R' {
			push(&rq, i)
		} else {
			push(&dq, i)
		}
	}
	for len(dq) > 0 && len(rq) > 0 {
		d0 := pop(&dq)
		r0 := pop(&rq)
		if d0 < r0 {
			push(&dq, d0+len(senate))
		} else {
			push(&rq, r0+len(senate))
		}
	}
	if len(dq) > 0 {
		return "Dire"
	}
	return "Radiant"
}
