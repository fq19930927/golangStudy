package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	count := 0
	for i := 0; i < 1000; i++ {
		go func() {
			count++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("count = %d", count)
}

//test Mutex
func TestCountThreadSafe(t *testing.T) {
	var mut sync.Mutex
	count := 0
	for i := 0; i < 1000; i++ {

		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
		}()
	}
	//不能删
	time.Sleep(time.Second * 1)
	t.Logf("count = %d", count)
}

//test WaitGroup
func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("count = %d", count)
}
