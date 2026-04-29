# Go 并发速查表

> Go 并发的核心概念、CSP 模型、同步原语一览。

---

## goroutine

**创建 goroutine：**
```go
go func() {
    fmt.Println("异步执行")
}()

go func(msg string) {
    fmt.Println(msg)
}("hello")
```

**Java 对比：**
- Java Thread 是重量级（1:1），goroutine 是轻量级（M:N）
- goroutine 启动比 Thread 快 100 倍
- Go 默认使用所有 CPU 核，不需要手动设置

---

## channel

**创建 channel：**
```go
ch := make(chan int)        // unbuffered
ch := make(chan int, 3)     // buffered，容量 3
```

**发送与接收：**
```go
ch <- 42        // 发送
v := <-ch      // 接收
<-ch           // 接收但不使用
```

**关闭 channel：**
```go
close(ch)
// 关闭后，接收操作不阻塞，返回零值
// 已关闭的 channel 不会再有数据
```

**遍历 channel：**
```go
for v := range ch {
    fmt.Println(v)
    if v == 0 { break } // 遇到 0 退出
}
```

---

## buffered vs unbuffered channel

| 类型 | 行为 | 何时使用 |
|------|------|----------|
| `make(chan T)` | 同步发送/接收 | 需要严格同步时 |
| `make(chan T, n)` | 非阻塞直到缓冲区满 | 解耦生产/消费 |

```go
// unbuffered：发送会阻塞，直到有人接收
ch := make(chan int)
ch <- 1 // 如果没有 receiver，这行会阻塞

// buffered：发送不阻塞，直到缓冲区满
ch := make(chan int, 3)
ch <- 1 // 立即返回
ch <- 2 // 立即返回
ch <- 3 // 立即返回
ch <- 4 // 阻塞，直到有人接收
```

---

## select

**多路复用 channel：**
```go
select {
case v := <-ch1:
    fmt.Println("从 ch1 收到:", v)
case ch2 <- 100:
    fmt.Println("发送到 ch2 成功")
case <-time.After(time.Second):
    fmt.Println("超时")
default:
    fmt.Println("没有任何 channel 就绪")
}
```

**特点：**
- 类似 switch，但 case 是 channel 操作
- 哪个 case 就绪就执行哪个
- 都不就绪时执行 default（可选）
- 多个就绪时随机选一个

---

## sync.WaitGroup

**等待一组 goroutine 完成：**
```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1) // 添加计数
    go func(id int) {
        defer wg.Done() // 完成时减计数
        fmt.Println("Worker", id, "完成")
    }(i)
}

wg.Wait() // 阻塞，直到计数归零
fmt.Println("所有 worker 完成")
```

**Java 对比：** 类似 `CountDownLatch`，但可以动态添加

---

## sync.Mutex

**互斥锁，保护共享资源：**
```go
var mu sync.Mutex
var counter int

mu.Lock()
counter++
mu.Unlock()
```

**defer 确保解锁：**
```go
mu.Lock()
defer mu.Unlock()
// ... 使用共享资源 ...
// 函数结束自动解锁
```

**Java 对比：** 类似 `synchronized` 或 `ReentrantLock`

---

## sync.RWMutex

**读写锁，读多写少时用：**
```go
var mu sync.RWMutex
var cache = make(map[string]string)

// 读：可以多个并发读
mu.RLock()
value := cache["key"]
mu.RUnlock()

// 写：独占，写时不能读
mu.Lock()
cache["key"] = "new value"
mu.Unlock()
```

**Java 对比：** 类似 `ReentrantReadWriteLock`

---

## sync.Once

**只执行一次：**
```go
var once sync.Once
var instance *Singleton

func GetInstance() *Singleton {
    once.Do(func() {
        fmt.Println("初始化只执行一次")
        instance = &Singleton{}
    })
    return instance
}
```

**Java 对比：** 类似双重检查锁定（DCL），但更简单

---

## sync/atomic

**原子操作，无锁并发：**
```go
var n int64

atomic.AddInt64(&n, 1)        // n++
atomic.LoadInt64(&n)          // 读取
atomic.StoreInt64(&n, 100)    // 写入
```

**适合场景：**
- 计数器
- 标志位
- 简单状态

---

## context

**创建和取消：**
```go
// 根 context
ctx := context.Background()

// 带取消
ctx, cancel := context.WithCancel(context.Background())
defer cancel() // 记得取消

// 带超时
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

// 带截止时间
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(ctx, deadline)
defer cancel()
```

**在 HTTP 请求中传递：**
```go
func handler(w http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    // 使用 ctx
}
```

**检查取消：**
```go
select {
case <-ctx.Done():
    fmt.Println("取消:", ctx.Err())
    return
default:
    // 继续执行
}
```

**传递值：**
```go
// 存
ctx := context.WithValue(ctx, "userID", 123)

// 取
userID := ctx.Value("userID").(int)
```

---

## Worker Pool

**标准模式：**
```go
func workerPool(jobCh <-chan Job, resultCh chan<- Result, numWorkers int) {
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobCh {
                result := process(job)
                resultCh <- result
            }
        }(i)
    }

    wg.Wait()    // 等待所有 worker 完成
    close(resultCh)
}
```

**Java 对比：** 类似 `ExecutorService` + `LinkedBlockingQueue`

---

## Fan-out / Fan-in

**Fan-out：多个 goroutine 从同一 channel 读取：**
```go
for i := 0; i < numWorkers; i++ {
    go func() {
        for job := range jobs {
            process(job)
        }
    }()
}
```

**Fan-in：合并多个 channel 到一个：**
```go
func merge(chs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    for _, ch := range chs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }

    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

---

## 常见死锁场景

```go
// 1. 发送和接收都在同一个 goroutine
ch := make(chan int)
ch <- 1           // 死锁：没有 receiver
<-ch              // 不会执行到这里

// 2. 没有人发送
ch := make(chan int)
<-ch              // 死锁：没有人发送

// 3. select 中所有 case 都阻塞且无 default
select {}         // 永久阻塞

// 4. 两 goroutine 互相等待
// goroutine A 等待 B 释放锁，B 等待 A 释放锁
```

---

## goroutine 泄漏

**泄漏原因：**
- channel 发送但无人接收
- select 阻塞但所有 case 都阻塞且无 default
- 忘记终止后台 goroutine

**避免方法：**
```go
// 使用 context 管理生命周期
ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return // 退出
        case job := <-jobs:
            process(job)
        }
    }
}()
```

---

## 并发安全 map

**sync.Map（适合读多写少）：**
```go
var m sync.Map

m.Store("key", "value")
value, ok := m.Load("key")
m.Delete("key")
m.Range(func(k, v any) bool {
    fmt.Println(k, v)
    return true
})
```

**普通 map + mutex：**
```go
var mu sync.RWMutex
var m = make(map[string]string)

mu.Lock()
m["key"] = "value"
mu.Unlock()

mu.RLock()
v := m["key"]
mu.RUnlock()
```

---

## 常见并发模式

**pipeline 模式：**
```go
// stage1 → stage2 → stage3
ch1 := gen(1, 2, 3)
ch2 := double(ch1)
ch3 := print(ch2)
```

**timeout 模式：**
```go
select {
case v := <-ch:
    fmt.Println(v)
case <-time.After(time.Second):
    fmt.Println("超时")
}
```

**quit channel 模式：**
```go
quit := make(chan struct{})
workers := make([]chan Job, 5)

for i := range workers {
    go worker(workers[i], quit)
}

// 通知所有 worker 退出
close(quit)
```

---

## 快速对照表

| 概念 | Go | Java |
|------|-----|------|
| 线程 | goroutine | Thread |
| 线程池 | goroutine + channel | ExecutorService |
| 队列 | channel | BlockingQueue |
| 锁 | sync.Mutex / RWMutex | synchronized / Lock |
| 计数器 | atomic / sync.WaitGroup | AtomicInteger / CountDownLatch |
| 线程本地 | 包级变量 + 传递 | ThreadLocal |
| 取消 | context | Future.cancel() |
| 超时 | context.WithTimeout | CompletableFuture.orTimeout() |