package limiter

import (
	"sync/atomic"
	"time"
)

// bucket
type bucket struct {
	ts    int64
	count int64
	_     [48]byte
}

// SlideWindowLimiter enforces a rate limit on operations using a sliding window algorithm.
type SlideWindowLimiter struct {
	interval int64 //间隔(毫秒)
	size     int64 //桶数量
	limit    int64 //窗口内最大请求数（limit/window_second=qps）
	total    int64 //当前窗口内的总请求数
	buckets  []bucket
}

func NewSlideWindowLimiter(limit int64, window time.Duration, bucketNum int64) *SlideWindowLimiter {
	interval := window.Milliseconds() / bucketNum
	if interval <= 0 {
		interval = 1
	}
	return &SlideWindowLimiter{
		limit:    limit,
		interval: interval,
		size:     bucketNum,
		total:    0,
		buckets:  make([]bucket, bucketNum),
	}
}
func (l *SlideWindowLimiter) Allow() bool {
	now := time.Now().UnixMilli()
	// 获取循环数组索引
	b := &l.buckets[(now/l.interval)%l.size]
	if ts := atomic.LoadInt64(&b.ts); now-ts >= l.interval { //清理过期bucket
		if atomic.CompareAndSwapInt64(&b.ts, ts, now) {
			if old := atomic.SwapInt64(&b.count, 0); old > 0 {
				atomic.AddInt64(&l.total, -old)
			}
		}
	}
	atomic.AddInt64(&b.count, 1)
	total := atomic.AddInt64(&l.total, 1)
	if total > l.limit {
		atomic.AddInt64(&b.count, -1)
		atomic.AddInt64(&l.total, -1)
		return false
	}
	return true
}
