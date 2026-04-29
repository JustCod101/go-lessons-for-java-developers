package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: received cancel signal, stopping...\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	// 1. WaitGroup: 等待一组 goroutine 完成
	// Java 对比: 类似于 CountDownLatch
	var wg sync.WaitGroup

	// 2. Context: 并发控制与超时
	// Java 对比: 类似于 Future.cancel() 或 ThreadLocal (用于传值)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 保证资源释放

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// 3. Mutex: 互斥锁
	// Java 对比: 类似于 ReentrantLock
	var mu sync.Mutex
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Final count: %d\n", count)
	fmt.Println("Main: all workers stopped or finished")
}
