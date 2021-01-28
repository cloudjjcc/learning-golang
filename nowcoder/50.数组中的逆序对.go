package main

import "fmt"

//题目描述
//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P%1000000007
//输入描述:
//题目保证输入的数组中没有的相同的数字
//数据范围：
//对于%50的数据,size<=10^4
//对于%75的数据,size<=10^5
//对于%100的数据,size<=2*10^5

func main() {
	testArr := []int{364, 637, 341, 406, 747, 995, 234, 971, 571, 219, 993, 407, 416, 366, 315, 301, 601, 650, 418, 355, 460, 505, 360, 965, 516, 648, 727, 667, 465, 849, 455, 181, 486, 149, 588, 233, 144, 174, 557, 67, 746, 550, 474, 162, 268, 142, 463, 221, 882, 576, 604, 739, 288, 569, 256, 936, 275, 401, 497, 82, 935, 983, 583, 523, 697, 478, 147, 795, 380, 973, 958, 115, 773, 870, 259, 655, 446, 863, 735, 784, 3, 671, 433, 630, 425, 930, 64, 266, 235, 187, 284, 665, 874, 80, 45, 848, 38, 811, 267, 575}
	fmt.Println(inversePairs(testArr))
}

// 归并排序的变形
func inversePairs(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	return mergeCore(arr, make([]int, len(arr))) % 1000000007
}

func mergeCore(arr, tmp []int) int {
	if len(arr) <= 1 {
		return 0
	}
	mid := len(arr) / 2
	leftCount := mergeCore(arr[:mid], tmp[:mid])
	rightCount := mergeCore(arr[mid:], tmp[mid:])
	// merge
	var (
		p1, p2 = mid - 1, len(arr) - 1
		i      = p2
		count  = 0
	)
	for p1 >= 0 && p2 >= mid {
		if arr[p1] > arr[p2] {
			tmp[i] = arr[p1]
			count += p2 - mid + 1
			p1--
		} else {
			tmp[i] = arr[p2]
			p2--
		}
		i--
	}
	// copy the rest
	for p1 >= 0 {
		tmp[i] = arr[p1]
		p1--
		i--
	}
	for p2 >= mid {
		tmp[i] = arr[p2]
		p2--
		i--
	}
	// copy tmp to arr
	for j := 0; j < len(arr); j++ {
		arr[j] = tmp[j]
	}
	return leftCount + rightCount + count
}
