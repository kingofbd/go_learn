package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func main() {
	wg := sync.WaitGroup{}
	counter := Counter{}

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
	fmt.Printf("最终计数器值: %d\n", counter.value) // 输出: 10000
}

// 定义计数器结构体
type Counter struct {
	mu    sync.Mutex
	value int
}

// 计数器自增
func (c *Counter) increment() {
	// 上锁
	c.mu.Lock()
	// 自增1
	c.value++
	// 释放锁
	c.mu.Unlock()
}
