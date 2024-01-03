package pool

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := NewPool(5)
	for i := 0; i < 10; i++ {
		count := i
		pool.AddJob(func() {
			println("exec ", count)
			time.Sleep(1 * time.Second)
		})
	}
	time.Sleep(3 * time.Second)
	pool.StopAll()
	pool.Wait()
	println("all finish")
}
