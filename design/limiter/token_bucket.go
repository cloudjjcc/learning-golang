package limiter

import (
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	Cap      int64
	Tokens   int64
	Rate     float64
	LastTime time.Time
	mu       sync.Mutex
}

func (l *TokenBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()

	// 生成tokens
	tokens := int64(now.Sub(l.LastTime).Seconds() * l.Rate)

	if tokens > 0 {
		l.Tokens += tokens
		if l.Tokens > l.Cap {
			l.Tokens = l.Cap
		}
		l.LastTime = now
	}

	// 获取tokens
	if l.Tokens <= 0 {
		return false
	}
	l.Tokens--
	return true
}
