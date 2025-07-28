package main

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func main() {
	// 创建同步等待组
	wg := sync.WaitGroup{}
	// 需要等待两个协程完成
	wg.Add(2)

	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 != 0 {
				fmt.Println("打印奇数的携程", i)
				time.Sleep(1 * time.Second)
			}
		}
		// 协程结束时标记完成
		defer wg.Done()
	}()

	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("打印偶数的携程", i)
				time.Sleep(1 * time.Second)
			}
		}
		// 协程结束时标记完成
		defer wg.Done()
	}()

	// 阻塞主协程直到所有协程完成
	wg.Wait()
}
