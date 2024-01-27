package huawei

import (
	"fmt"
	"sort"
)

func mainBossIncome() {
	count := 0
	if _, err := fmt.Scanln(&count); err != nil {
		return
	}
	records := make([]*record, 0, count)
	for {
		var r record
		_, err := fmt.Scanln(&r.id, &r.pid, &r.amount)
		if err != nil {
			break
		}
		records = append(records, &r)
		if len(records) == count {
			break
		}
	}
	bossId, bossAmount := bossIncome(records)
	fmt.Println(bossId, bossAmount)
}

type record struct {
	id     int
	pid    int
	amount int
}

func bossIncome(records []*record) (int, int) {
	sort.Slice(records, func(i, j int) bool {
		if records[i].pid > records[j].pid {
			return true
		}
		return false
	})
	bossId := records[len(records)-1].pid
	m := make(map[int]int)
	for _, v := range records {
		v.amount += m[v.id]
		deductAmount := 15 * (v.amount / 100)
		v.amount -= deductAmount
		m[v.pid] += deductAmount
	}
	return bossId, m[bossId]
}
