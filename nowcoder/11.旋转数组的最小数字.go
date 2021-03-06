package nowcoder

//题目描述
//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
//例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
//NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

// O(n)
func minNumberInRotateArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	preNum := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < preNum {
			return arr[i]
		}
	}
	return arr[0]
}

// binary search O(log n)
func minNumberInRotateArray2(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var (
		start = 0
		mid   = 0
		end   = len(arr) - 1
	)
	for start+1 < end {
		if arr[start] < arr[end] {
			return arr[start]
		}
		mid = (start + end) / 2
		if arr[mid] < arr[start] {
			end = mid
		} else if arr[mid] > arr[start] {
			start = mid
		} else {
			start++
		}
	}
	return arr[end]
}
