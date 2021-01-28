package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

//题目描述
//将一个字符串转换成一个整数，要求不能使用字符串转换整数的库函数。
//数值为0或者字符串不是一个合法的数值则返回0

func main() {
	testStr := "-111111111111111111111111"
	fmt.Println(myitoa(testStr))
	fmt.Println(strconv.Atoi(testStr))
}

func myitoa(str string) (r int, err error) {
	// empty string
	if len(str) == 0 {
		err = errors.New("invalid input")
		return
	}
	// '+' or '-'
	flag := 1
	if str[0] == '-' {
		flag = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	for _, v := range str {
		if v > '9' || v < '0' {
			err = errors.New("invalid input")
			return 0, err
		}
		r *= 10
		r += int(v - '0')
	}
	if r > math.MaxInt32 {
		err = errors.New("overflow")
		return 0, err
	}
	return flag * r, nil
}
