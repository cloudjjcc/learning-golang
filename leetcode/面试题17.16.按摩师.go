package leetcode

//一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。
//在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。
//给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
//注意：本题相对原题稍作改动
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/the-masseuse-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// f(x)=max(f[x-1],f[x-2]+arr[x])
func massage(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	pre1 := 0
	pre2 := 0
	for i := 0; i < len(arr); i++ {
		max := pre1
		if max < pre2+arr[i] {
			max = pre2 + arr[i]
		}
		pre2 = pre1
		pre1 = max
	}
	return pre1
}
