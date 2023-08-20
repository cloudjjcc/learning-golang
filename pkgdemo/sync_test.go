package pkgdemo

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestMutex(t *testing.T) {
	mutex := sync.Mutex{}
	mutex.Lock()
	//mutex.TryLock()
	mutex.Unlock()
	t.Logf("%v", &mutex)
}
func TestRWMutex(t *testing.T) {
	rw := sync.RWMutex{}
	for i := 0; i < 10; i++ {
		rw.RLock()
	}
	rw.RUnlock()
	rw.Lock()
}
func TestAtomic(t *testing.T) {
	var arr []int
	value := atomic.Value{}
	value.Store(arr)
	t.Logf("%v", value.Load())
}

func TestPool(t *testing.T) {

	pool := sync.Pool{
		New: func() interface{} {
			return make([]int, 0)
		},
	}

	for i := 0; i < 100; i++ {
		get := pool.Get().([]int)
		pool.Put(get)
	}
	t.Logf("%v", &pool)
}

func TestMap(t *testing.T) {
	m := sync.Map{}
	m.Store("a", "av")
	v, _ := m.Load("a")
	assert.Equal(t, "av", v.(string))
	runtime.GC()
}
