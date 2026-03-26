package pkgdemo

import (
	"fmt"
	"sync"
)

func DemoErrGroup() {
	sm := sync.Map{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			sm.Store(i, i)
		}()
	}
	wg.Wait()
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	fmt.Println("done")
	pool := sync.Pool{}
	pool.New = func() any {
		return 1
	}
}
