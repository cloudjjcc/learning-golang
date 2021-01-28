package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

//高并发下的锁与map的读写
//场景：在一个高并发的web服务器中，要限制IP的频繁访问。
//现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。
//每个IP三分钟之内只能访问一次。
//修改以下代码完成该过程，要求能成功输出 success:100
type Ban struct {
	visitIPs map[string]time.Time
	lock     sync.RWMutex
}

func NewBan(ctx context.Context) *Ban {
	ban := &Ban{visitIPs: make(map[string]time.Time)}
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for {
			select {
			case <-ticker.C:
				ban.lock.Lock()
				for ip, t := range ban.visitIPs {
					if time.Now().Sub(t) > time.Minute*3 {
						delete(ban.visitIPs, ip)
					}
				}
				ban.lock.Unlock()
			case <-ctx.Done():
				return
			}
		}
	}()
	return ban
}
func (o *Ban) visit(ip string) bool {
	o.lock.Lock()
	defer o.lock.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := int32(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ban := NewBan(ctx)
	wg := sync.WaitGroup{}
	wg.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(ii int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", ii)
				if !ban.visit(ip) {
					atomic.AddInt32(&success, 1)
				}
			}(j)
		}

	}
	wg.Wait()
	fmt.Println("success:", success)
}
