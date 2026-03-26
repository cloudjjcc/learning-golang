package nowcoder

//题目描述
//在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P%1000000007
//输入描述:
//题目保证输入的数组中没有的相同的数字
//数据范围：
//对于%50的数据,size<=10^4
//对于%75的数据,size<=10^5
//对于%100的数据,size<=2*10^5

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
