# 进阶练习

本目录包含 Go 语言进阶练习，适合学完前 13 节课程的学习者。

---

## 练习清单

### 1. 并发爬虫
**目标**: 实现并发网页抓取
**输入输出示例**:
```
输入: ["https://example.com", "https://go.dev"]
输出:
- https://example.com: 标题 "Example Domain"
- https://go.dev: 标题 "The Go Programming Language"
```
**提示**: 使用 goroutine + channel + worker pool，并发抓取多个 URL
**参考解法路径**: 创建 URL channel，启动 N 个 worker 并发抓取，使用 sync.WaitGroup 等待完成

---

### 2. 限流器
**目标**: 实现令牌桶限流器
**输入输出示例**:
```
限制: 每秒 2 个请求
发送 5 个请求...
请求 1: 通过 (time: 0s)
请求 2: 通过 (time: 0.5s)
请求 3: 被限流 (time: 1s)
请求 4: 通过 (time: 1.5s)
请求 5: 通过 (time: 2s)
```
**提示**: 用 channel + select 实现限流
**参考解法路径**:
```go
type RateLimiter struct {
    tokens chan struct{}
    ticker *time.Ticker
}
```

---

### 3. TTL 缓存
**目标**: 实现带过期时间的缓存
**输入输出示例**:
```
Set("key1", "value1", 2s)
Get("key1"): value1 (2秒内)
睡眠 3 秒
Get("key1"): nil, false (已过期)
```
**提示**: 存储时记录过期时间，获取时检查是否过期，用 goroutine 定期清理
**参考解法路径**: `map[string]struct{value any; expireAt time.Time}`，用 `time.Now().After(expireAt)` 判断

---

### 4. 连接池模拟
**目标**: 实现简单的数据库连接池
**输入输出示例**:
```
获取连接: conn-1
获取连接: conn-2
获取连接: conn-3
连接池已满，等待...
释放连接: conn-1
获取连接: conn-1
```
**提示**: 预先创建 N 个连接，用 channel 管理借出和归还
**参考解法路径**: `make(chan *Connection, poolSize)`，获取时 `<-pool`，归还时 `pool <- conn`

---

### 5. 简单 RPC
**目标**: 实现简单的 RPC 调用
**输入输出示例**:
```
服务端注册: Add
客户端调用: Add(3, 5) = 8
```
**提示**: 用 net/http 实现，利用 `context.WithValue` 传递方法名和参数
**参考解法路径**: 服务端用 Handler 接收请求，解析方法名和参数，调用本地函数返回结果

---

### 6. Graceful Shutdown
**目标**: 实现优雅关闭
**输入输出示例**:
```
启动服务器 :8080
收到 SIGTERM
停止接受新请求
处理进行中的请求 (3个)
服务器关闭
```
**提示**: 用 `os/signal` 捕获信号，停止接受新请求，等待进行中的请求完成
**参考解法路径**:
```go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, syscall.SIGTERM)
<-sigCh
// 停止监听，关闭服务器
```

---

### 7. Context 超时控制
**目标**: 实现带超时的操作
**输入输出示例**:
```
操作启动
操作完成 (耗时: 1.5s, 无超时)
---
操作启动
超时 (耗时: 2s)
```
**提示**: 用 `context.WithTimeout` 创建带超时的 context，超时后操作被取消
**参考解法路径**:
```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
select {
case <-doLongOperation():
    fmt.Println("操作完成")
case <-ctx.Done():
    fmt.Println("超时")
}
```

---

### 8. 日志中间件
**目标**: 实现 HTTP 日志中间件
**输入输出示例**:
```
GET /hello 200 OK (耗时: 1.234ms)
POST /api/data 201 Created (耗时: 56.789ms)
GET /notfound 404 Not Found (耗时: 0.567ms)
```
**提示**: 中间件是一个函数，接收 `http.Handler` 返回 `http.Handler`
**参考解法路径**:
```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        fmt.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
    })
}
```

---

### 9. 简单任务调度器
**目标**: 实现定时任务调度器
**输入输出示例**:
```
启动调度器
任务 A 将在 1 秒后执行
任务 B 将在 3 秒后执行
任务 A 执行 (time: 1s)
任务 C 将在 5 秒后执行
任务 B 执行 (time: 3s)
任务 C 执行 (time: 5s)
调度器关闭
```
**提示**: 用 `time.Timer` 或 `time.Ticker` 实现定时任务
**参考解法路径**: `type Task struct { Run func(); Next time.Time }`，用优先队列或定时器管理

---

### 10. 上下文传递
**目标**: 在 goroutine 间传递 context
**输入输出示例**:
```
[主goroutine] 创建带超时的 context
[Worker 1] 开始工作
[Worker 2] 开始工作
[Worker 1] 完成
[主goroutine] 被取消，超时
[Worker 2] 被中断
```
**提示**: context 在 goroutine 间传递，父 context 取消时子 context 也会被取消
**参考解法路径**: `ctx, cancel := context.WithTimeout(parentCtx, duration)`，传递给子 goroutine

---

### 11. 读写锁保护缓存
**目标**: 用 RWMutex 实现高效缓存
**输入输出示例**:
```
Read: value1 (耗时: 200ns)
Write: 更新 value1
Read: value2 (耗时: 180ns)
```
**提示**: 读多写少时用 RWMutex，允许多个读锁同时存在，写锁独占
**参考解法路径**:
```go
var mu sync.RWMutex
func Read(key string) any {
    mu.RLock()
    defer mu.RUnlock()
    return cache[key]
}
func Write(key string, value any) {
    mu.Lock()
    defer mu.Unlock()
    cache[key] = value
}
```

---

### 12. Once 单例初始化
**目标**: 用 sync.Once 实现单例
**输入输出示例**:
```
GetInstance() 第1次调用: 初始化完成
GetInstance() 第2次调用: 返回已存在的实例
GetInstance() 第3次调用: 返回已存在的实例
```
**提示**: `sync.Once` 保证初始化代码只执行一次，线程安全
**参考解法路径**:
```go
var once sync.Once
var instance *Singleton
func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

---

### 13. Atomic 计数器
**目标**: 用 atomic 包实现无锁计数器
**输入输出示例**:
```
初始值: 0
Add 5: 5
Add 3: 8
Increment: 9
Load: 9
```
**提示**: `sync/atomic` 提供原子操作，比 mutex 轻量
**参考解法路径**: `atomic.AddInt64(&counter, 5)`，`atomic.LoadInt64(&counter)`

---

### 14. 多路复用 select 与超时
**目标**: 综合使用 select 和 channel
**输入输出示例**:
```
同时监听 3 个 channel
收到 ch1: 100
收到 ch2: 200
超时 (等待 500ms)
```
**提示**: select 监听多个 channel，可以有 default case 和 timeout case
**参考解法路径**:
```go
select {
case v := <-ch1:
    fmt.Println("收到 ch1:", v)
case v := <-ch2:
    fmt.Println("收到 ch2:", v)
case <-time.After(500 * time.Millisecond):
    fmt.Println("超时")
}
```

---

### 15. 组合构建 HTTP 中间件链
**目标**: 实现中间件的函数式组合
**输入输出示例**:
```
中间件链: Logger → Recover → Auth
请求: /api/data (带 token)
[Logger] 开始
[Recover] 开始
[Auth] 验证通过
[Auth] 结束
[Recover] 结束
[Logger] 结束
响应: 200 OK
```
**提示**: 中间件是返回 `http.Handler` 的函数，可以用函数组合简化调用
**参考解法路径**:
```go
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}
```

---

## 学习完成

完成进阶练习后，你可以进入 [projects/](../projects/README.md) 开始实战项目训练。