package limiter

import (
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	Cap      int64     //容量
	Water    int64     //水位(存储请求)
	Rate     float64   //每秒出水量
	LastTime time.Time //最后一次访问时间
	mu       sync.Mutex
}

func (l *LeakyBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	//计算漏水
	leaked := int64(now.Sub(l.LastTime).Seconds() * l.Rate)
	if leaked > 0 {
		l.Water -= leaked
		if l.Water <= 0 {
			l.Water = 0
		}
		l.LastTime = now
	}
	// 桶满
	if l.Water >= l.Cap {
		return false
	}
	// 桶里水量增加
	l.Water++
	return true
}
