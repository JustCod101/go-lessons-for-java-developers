---
title: "Lesson 15: 综合项目 - 内存缓存"
chapter: 15
part: Go高级特性
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 内存缓存
prerequisites:
  - "[[lesson-14/README]]"
related:
  - "[[lesson-14/README]]"
  - "[[lesson-16/README]]"
key_concepts:
  - map
  - goroutine
  - channel
  - sync.RWMutex
  - TTL
  - Benchmark
---

# Lesson 15: 综合项目 - 内存缓存

欢迎来到第 15 课。今天我们不讲新语法，而是把之前学到的 `map`、`goroutine`、`channel` 和 `sync` 包组合起来，亲手写一个高性能的内存缓存。

### 1. 学习目标
* 掌握如何使用 `sync.RWMutex` 保护共享资源。
* 理解缓存过期清理（TTL）的实现机制。
* 学习如何在 Go 中编写简单的 [[Benchmark]] 测试。
* 能够对比 Go 缓存实现与 [[Java]] `ConcurrentHashMap` 的异同。

### 2. 给 [[Java开发者|Java 开发者]]的类比
在 Java 中，如果你要实现一个带过期的本地缓存，通常会用 `ConcurrentHashMap` 存储数据，再配合一个 `ScheduledExecutorService` 定期清理过期 Key。或者你会直接用 Guava Cache 或 Caffeine。

在 Go 中，我们没有内置的 `ConcurrentHashMap`（虽然有 `sync.Map`，但它有特定的适用场景）。我们通常使用 `map` + `sync.RWMutex` 来实现。清理机制则通过一个后台 [[Goroutine|goroutine]] 配合 `time.Ticker` 来完成。

### 3. 核心概念

#### 并发安全：[[RWMutex]] vs Mutex
**[[Java对比|Java 对比]]**：类似于 `ReentrantReadWriteLock`。
缓存通常是“读多写少”的场景。使用 `sync.RWMutex` 可以让多个 goroutine 同时读取，只有写入时才会互斥，这比单纯的 `sync.Mutex` 性能更高。

#### 过期机制 (TTL)
**Java 对比**：类似于 `Caffeine` 的 `expireAfterWrite`。
我们需要在存储数据时记录一个过期时间点。后台启动一个“看门狗” goroutine，每隔一段时间扫描一遍 [[Map|map]]，删掉那些已经过期的条目。

#### 性能测试 (Benchmark)
**Java 对比**：类似于 JMH (Java Microbenchmark Harness)。
Go 原生支持基准测试。只需在 `_test.go` 文件中编写以 `Benchmark` 开头的函数，并使用 `go test -bench=.` 即可运行。

### 4. 代码示例
详细实现请参考 `src/cache.go`。这里展示核心结构：

```go
type Item struct {
    Value      interface{}
    Expiration int64 // 过期时间戳
}

type Cache struct {
    items map[string]Item
    mu    sync.RWMutex
}

// Get 读取数据，使用 RLock
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    
    item, found := c.items[key]
    if !found || (item.Expiration > 0 && time.Now().UnixNano() > item.Expiration) {
        return nil, false
    }
    return item.Value, true
}
```

### 5. 常见误区
1. **忘记释放锁**：在 Java 中我们有 `try-finally`，Go 中一定要记得 `defer mu.Unlock()`。
2. **在循环中直接删除 map 元素**：Go 的 map 在迭代时删除元素是安全的，但如果你在迭代时没有加写锁，而另一个 goroutine 正在读取，会触发 `fatal error: concurrent map iteration and map write`。
3. **TTL 精度问题**：后台清理任务不宜频率过高（浪费 CPU）或过低（浪费内存）。

### 6. 本节练习
1. **实现清空功能**：为 `Cache` 增加一个 `Flush()` 方法，清空所有数据。
2. **增加命中率统计**：统计缓存的 Hit 和 Miss 次数。
3. **优化清理逻辑**：目前的清理是全量扫描，思考如何在大数据量下优化（提示：分片锁或随机采样）。

### 7. 面试可能怎么问
* **Q: 为什么不用 [[sync.Map]] 实现缓存？**
  * A: `sync.Map` 适用于 key 不怎么变，或者多个 goroutine 读写不相交的场景。对于频繁写入和删除的缓存，`map + RWMutex` 通常性能更好且更灵活。
* **Q: Go 的 map 为什么不是 [[并发安全]] 的？**
  * A: 为了性能。大多数场景下 map 不需要并发保护，强制加锁会拖慢速度。Go 选择了把控制权交给开发者。
* **Q: 如何测试缓存的并发性能？**
  * A: 使用 `testing` 包的 `Benchmark` 功能，配合 `b.RunParallel` 模拟多核并发。

### 8. 本节总结
通过这个小项目，你应该能感觉到 Go 处理并发的直观性。我们不需要复杂的 [[ExecutorService|线程池]] 管理，一个 `go func()` 就能解决后台任务。记住：**不要通过共享内存来通信，而要通过通信来共享内存**（虽然在缓存这种底层组件中，锁依然是必不可少的）。
