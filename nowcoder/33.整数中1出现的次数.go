package main

import "fmt"

//题目描述
//求出1~13的整数中1出现的次数,并算出100~1300的整数中1出现的次数？
//为此他特别数了一下1~13中包含1的数字有1、10、11、12、13因此共出现6次,
//但是对于后面问题他就没辙了。
//ACMer希望你们帮帮他,并把问题更加普遍化,
//可以很快的求出任意非负整数区间中1出现的次数（从1 到 n 中1出现的次数）。

func main() {
	fmt.Println(numberOf1BetweenN(10000))
}

// 数学归纳法
// 个位： (n/10)*1+(n%10)>=2*1?1:(n%10-1)+1
// 十位： (n/100)*10+(n%100)>=2*10?10:(n%100-10)+1
// 百位： (n/1000)*100+(n%1000)>=200?100:(n%1000-100)+1
// 总结： (n/i*10)*i+(n%i*10)>=2*i?i:(n%i*10-i)+1
func numberOf1BetweenN(n int) int {
	if n <= 0 {
		return 0
	}
	count := 0
	for i := 1; i <= n; i *= 10 {
		a := i * 10
		b := n / a
		c := n % a
		sum := b * i
		if c >= 2*i {
			sum += i
		} else if c >= i {
			sum += c - i + 1
		}
		count += sum
	}
	return count
}
