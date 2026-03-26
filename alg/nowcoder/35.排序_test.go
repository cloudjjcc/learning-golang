package nowcoder

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_sort(t *testing.T) {
	count := int(1e6)
	testArr := make([]int, count)
	rand.Seed(time.Now().Unix())
	for i := 0; i < count-1; i++ {
		testArr[i] = rand.Intn(10000)
	}
	start := time.Now()
	//BubbleSort(testArr)
	//InsertionSort(testArr)
	//ShellSort(testArr)
	//SelectionSort(testArr)
	//HeapSort(testArr)
	//MergeSort(testArr)
	QuickSort(testArr)
	//sort.Sort(sortArr(testArr))
	fmt.Println(time.Since(start))
}
