package main

import (
	"fmt"
	"sync"
	"time"
)

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func main() {
	// 初始化任务列表
	tasks := make([]Task, 0)
	tasks = append(tasks, Task{"任务1", func() {
		time.Sleep(2 * time.Second)
	}})
	tasks = append(tasks, Task{"任务2", func() {
		time.Sleep(1 * time.Second)
	}})
	tasks = append(tasks, Task{"任务3", func() {
		time.Sleep(3 * time.Second)
	}})

	// 执行任务
	wg := sync.WaitGroup{}
	do(tasks, &wg)
	/**
	[任务2] 执行完成 | 耗时: 1.001077375s
	[任务1] 执行完成 | 耗时: 2.001075333s
	[任务3] 执行完成 | 耗时: 3.001084083s
	*/
	wg.Wait()
}

// 任务结构体
type Task struct {
	// 任务编号
	id string
	// 任务具体执行的函数
	Job func()
}

// 执行任务
func do(tasks []Task, wg *sync.WaitGroup) {
	// 遍历任务列表，使用协程并发开启
	for _, task := range tasks {
		wg.Add(1)
		go func(task Task) {
			// 记录任务开启时间
			taskStart := time.Now()
			task.Job()
			duration := time.Since(taskStart)
			fmt.Printf("[%s] 执行完成 | 耗时: %v\n", task.id, duration)
			defer wg.Done()
		}(task)
	}
}
