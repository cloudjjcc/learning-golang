package nowcoder

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

//题目描述
//输入一个正整数数组，把数组里所有数字拼接起来排成一个数，
//打印能拼接出的所有数字中最小的一个。例如输入数组{3，32，321}，
//则打印出这三个数字能排成的最小数字为321323。

type MyInt []int

func (m MyInt) Len() int {
	return len(m)
}

func (m MyInt) Less(i, j int) bool {
	return fmt.Sprintf("%d%d", i, j) > fmt.Sprintf("%d%d", j, i)
}

func (m MyInt) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func PrintMinNum(arr []int) string {
	if len(arr) == 0 {
		return ""
	}
	// sort
	sort.Sort(MyInt(arr))
	// arrange
	buf := bytes.NewBuffer(make([]byte, 10))
	for i := 0; i < len(arr); i++ {
		buf.WriteString(strconv.Itoa(arr[i]))
	}
	return buf.String()
}
