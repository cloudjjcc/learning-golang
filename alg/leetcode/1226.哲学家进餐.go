package leetcode

import (
	"fmt"
	"go.uber.org/atomic"
	"time"
)

//5 个沉默寡言的哲学家围坐在圆桌前，每人面前一盘意面。
//叉子放在哲学家之间的桌面上。（5 个哲学家，5 根叉子）
//所有的哲学家都只会在思考和进餐两种行为间交替。
//哲学家只有同时拿到左边和右边的叉子才能吃到面，而同一根叉子在同一时间只能被一个哲学家使用。
//每个哲学家吃完面后都需要把叉子放回桌面以供其他哲学家吃面。
//只要条件允许，哲学家可以拿起左边或者右边的叉子，但在没有同时拿到左右叉子时不能进食。
//假设面的数量没有限制，哲学家也能随便吃，不需要考虑吃不吃得下。
//设计一个进餐规则（并行算法）使得每个哲学家都不会挨饿；
//也就是说，在没有人知道别人什么时候想吃东西或思考的情况下，每个哲学家都可以在吃饭和思考之间一直交替下去。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/the-dining-philosophers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

var forks [5]atomic.Int32
var pCh = make(chan struct{}, 4)

// 拿叉子
func pickFork(p, idx int) {
	for !forks[idx].CAS(0, 1) {
		time.Sleep(1 * time.Microsecond)
	}
	fmt.Printf("哲学家 %d pick %d fork\n", p, idx)
}

// 放下叉子
func putFork(p, idx int) {
	forks[idx].Store(0)
	fmt.Printf("哲学家 %d put %d fork\n", p, idx)
}

// 吃面
func eat(p int) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("哲学家%d开始吃面\n", p)
}
func wantsToEat(philosopher int) {
	for {
		left := philosopher
		right := (philosopher + 1) % 5
		pCh <- struct{}{}
		pickFork(philosopher, left)
		pickFork(philosopher, right)
		eat(philosopher)
		putFork(philosopher, left)
		putFork(philosopher, right)
		<-pCh
	}
}
