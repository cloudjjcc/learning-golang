package limiter

import (
	"sync"
	"time"
)

type Limiter interface {
	Allow() bool
}
type CounterLimiter struct {
	cnt        int
	limit      int
	windowSize int64
	start      int64
	lock       sync.Mutex
}

func (l *CounterLimiter) Allow() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	if time.Now().UnixNano()/1e6 > l.windowSize {
		l.cnt = 1
		l.start = time.Now().UnixNano() / 1e6
		return true
	}
	if l.cnt+1 <= l.limit {
		l.cnt++
		return true
	}
	return false
}
