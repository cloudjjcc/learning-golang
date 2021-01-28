package main

import (
	"fmt"
	"math"
)

//机器人坐标问题
//问题描述
//有一个机器人，给一串指令，L左转 R右转，F前进一步，B后退一步，
//问最后机器人的坐标，最开始，机器人位于 0 0，方向为正Y。
//可以输入重复指令n ： 比如 R2(LF) 这个等于指令 RLFLF。
//问最后机器人的坐标是多少？
type Vector struct {
	X float64
	Y float64
}

// 向量旋转
func (v *Vector) rotate(radians float64) {
	x0 := v.X
	y0 := v.Y
	v.X = math.Ceil(x0*math.Cos(radians) - y0*math.Sin(radians))
	v.Y = math.Ceil(x0*math.Sin(radians) + y0*math.Cos(radians))
}

// 向量相加
func (v *Vector) add(delta *Vector) {
	v.X += delta.X
	v.Y += delta.Y
}

// 向量相减
func (v *Vector) sub(delta *Vector) {
	v.X -= delta.X
	v.Y -= delta.Y
}

func resovleCmd(cmd string, pos *Vector, dir *Vector) {
	for _, r := range cmd {
		switch r {
		case 'L':
			dir.rotate(math.Pi / 2)
		case 'R':
			dir.rotate(-math.Pi / 2)
		case 'F':
			pos.add(dir)
		case 'B':
			pos.sub(dir)
		default:
			fmt.Println("error cmd :", r)
		}
	}

}
func main() {
	initPos := Vector{0, 0}
	initDir := Vector{0, 1}
	resovleCmd("LFF", &initPos, &initDir)
	fmt.Printf("final pos is:%+v,final dir is %+v\n", initPos, initDir)
}
