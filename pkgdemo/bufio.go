package pkgdemo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func demoScanner() {
	// 人为输入源。
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 设置扫描操作的分割功能。
	scanner.Split(bufio.ScanWords)
	// 计算单词。
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
}
