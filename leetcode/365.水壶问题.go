package leetcode

//有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？
//如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
//你允许：
//装满任意一个水壶
//清空任意一个水壶
//从一个水壶向另外一个水壶倒水，直到装满或者倒空
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/water-and-jug-problem
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 数学法：贝祖定理
// ax+by=z有解当且仅当z是x和y的最大公约数的倍数
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x, y) == 0
}
func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
