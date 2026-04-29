package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Millisecond * 500) // 模拟耗时任务
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	// 1. Goroutine: 轻量级线程
	// Java 对比: 类似于虚拟线程 (Project Loom)
	go func() {
		fmt.Println("Hello from goroutine")
	}()

	// 2. Channel: 用于通信
	// Java 对比: 类似于 BlockingQueue
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3. Worker Pool 模式
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 channel，通知 worker 任务结束

	// 收集结果
	for a := 1; a <= 5; a++ {
		res := <-results
		fmt.Printf("Result: %d\n", res)
	}

	// 4. Select: 多路复用
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
		}
	}
}
