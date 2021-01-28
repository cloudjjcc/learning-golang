package main

import "fmt"

//题目描述
//有一个XxY的网格，一个机器人只能走格点且只能向右或向下走，要从左上角走到右下角。
//请设计一个算法，计算机器人有多少种走法。
//注意这次的网格中有些障碍点是不能走的。
//给定一个int[][] map(C++ 中为vector >),表示网格图，
//若map[i][j]为1则说明该点不是障碍点，否则则为障碍。
//另外给定int x,int y，表示网格的大小。
//请返回机器人从(0,0)走到(x - 1,y - 1)的走法数，为了防止溢出，请将结果Mod 1000000007。
//保证x和y均小于等于50
func main() {
	testmarix := [][]int{
		{},
		{},
		{},
	}
	fmt.Println(countWay(testmarix, 50, 50))
}

// dp
// f(x,y)=f(x-1,y)+f(x,y-1)
func countWay(matrix [][]int, x, y int) int {
	// init x*y matrix
	f := make([][]int, x)
	for i := 0; i < x; i++ {
		f[i] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// 考虑障碍点
			//if matrix[i][j] == 1 {
			//	f[i][j] = 0
			//	continue
			//}
			if i == 0 && j == 0 {
				f[i][j] = 1
				continue
			}
			if i == 0 {
				f[i][j] = f[0][j-1]
				continue
			}
			if j == 0 {
				f[i][j] = f[x-1][0]
				continue
			}
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[x-1][y-1]
}
