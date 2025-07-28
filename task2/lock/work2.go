package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	counter := Counter2{}

	// 开启10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 1000; i++ {
				counter.increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("最终的计数器值: %d\n", counter.value) // 输出: 10000
}

// 无锁计数器结构体
type Counter2 struct {
	value int64
}

// 计数器自增
func (c *Counter2) increment() {
	atomic.AddInt64(&c.value, 1)
}
