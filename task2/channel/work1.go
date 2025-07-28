package main

import (
	"fmt"
	"sync"
	"time"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func main() {
	// 构建通道
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// 发送数据的协程
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 500)
		}
		// 数据发送完毕之后关闭通道
		close(ch)
		defer wg.Done()
	}()

	// 接收数据的协程
	go func() {
		// 检测到通道关闭后会退出for循环
		for num := range ch {
			fmt.Println(num)
		}
		defer wg.Done()
	}()

	wg.Wait()
}
