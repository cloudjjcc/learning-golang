package main

//题目描述
//需要保证奇数和奇数，偶数和偶数之间的相对位置不变，这和书本不太一样。

func main() {
	testArr := []int{1, 2, 3, 4, 5, 6}
	reOrderEvenOdd(testArr)
}

func reOrderEvenOdd(arr []int) {
	odd := make([]int, 0)
	even := make([]int, 0)

	for _, v := range arr {
		if v&1 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	for i, v := range even {
		arr[i] = v
	}
	for i, v := range odd {
		arr[len(even)+i] = v
	}
}

// 双指针
func reOrderEvenOdd2(arr []int) {
	if len(arr) <= 1 {
		return
	}

}
