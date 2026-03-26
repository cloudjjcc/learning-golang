package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main3() {
	scanner := bufio.NewScanner(os.Stdin)
	var line1 string
	if scanner.Scan() {
		line1 = scanner.Text()
	}
	numStrs := strings.Split(line1, " ")
	type job struct {
		submitAt int
		cost     int
	}
	jobs := make([]*job, 0, len(numStrs)/2)
	for i := 0; i < len(numStrs); i += 2 {
		t := &job{}
		t.submitAt, _ = strconv.Atoi(numStrs[i])
		t.cost, _ = strconv.Atoi(numStrs[i+1])
		jobs = append(jobs, t)
	}
	var queueMaxLen, count int
	if _, err := fmt.Scanln(&queueMaxLen, &count); err != nil {
		return
	}
	queue := make([]*job, 0, queueMaxLen)
	push := func(t *job) {
		queue = append(queue, t)
	}
	pop := func() *job {
		if len(queue) == 0 {
			return nil
		}
		t := queue[0]
		queue = queue[1:]
		return t
	}
	full := func() bool {
		return len(queue) == queueMaxLen
	}
	empty := func() bool {
		return len(queue) == 0
	}
	works := make([]int, count)
	totalTime := 0
	jobIdx := 0
	dropCount := 0
	for {
		if jobIdx >= len(jobs) {
			break
		}
		// 判断当前是否有空闲执行者
		idleWorkerIdx := -1
		for i := range works {
			if works[i] != 0 {
				works[i]--
			}
			if idleWorkerIdx == -1 && works[i] == 0 {
				idleWorkerIdx = i
			}
		}
		// 取出任务
		if idleWorkerIdx != -1 && !empty() {
			works[idleWorkerIdx] = pop().cost
		}
		// 有任务正在提交
		if totalTime == jobs[jobIdx].submitAt {
			if full() {
				pop()
				push(jobs[jobIdx])
				dropCount++
			} else {
				push(jobs[jobIdx])
			}
			jobIdx++
		}
		totalTime++
	}
	fmt.Printf("%d %d\n", totalTime+jobs[jobIdx-1].cost, dropCount)
}
