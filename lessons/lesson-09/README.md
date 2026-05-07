---
title: "Lesson 09: context 与并发控制"
chapter: 09
part: Go并发编程
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 并发控制
prerequisites:
  - "[[lesson-08/README]]"
related:
  - "[[lesson-08/README]]"
  - "[[lesson-10/README]]"
key_concepts:
  - context
  - WaitGroup
  - Mutex
  - RWMutex
  - atomic
  - ThreadLocal
---

# LESSON 09: [[context]] 与[[并发控制]]

在复杂的分布式系统中，如何优雅地取消一个请求或设置超时时间？Go 提供了 `context` 包来解决这个问题。同时，我们也会学习传统的同步原语。

## 1. 学习目标
* 理解 `context` 的核心作用：取消信号、超时控制和元数据传递。
* 掌握 `context.WithCancel`, `WithTimeout`, `WithDeadline` 的用法。
* 熟练使用 `sync.WaitGroup` 协调多个 [[Goroutine|goroutine]]。
* 掌握 `sync.Mutex` 和 `sync.RWMutex` 处理竞态条件。
* 了解 `sync.Once` 和 `atomic` 包的高级用法。

## 2. 给 [[Java开发者|Java 开发者]]的类比
* **context.Context**: 
    * 取消功能类似于 `Future.cancel()`。
    * 传值功能类似于 `ThreadLocal`，但它是显式传递的，更符合 Go 的显式哲学。
* **[[WaitGroup|sync.WaitGroup]]**: 类似于 `CountDownLatch`。
* **[[Mutex|sync.Mutex]]**: 类似于 `ReentrantLock`。
* **[[atomic]]**: 类似于 `java.util.concurrent.atomic` 包。

## 3. 核心概念

### Context (上下文)
`context` 是 Go [[并发编程]]的灵魂。它在 API 边界之间传递，用于控制生命周期。
* **取消信号**: 当父 context 被取消时，所有派生的子 context 都会收到 `Done()` 信号。
* **超时控制**: 自动在指定时间后发送取消信号。
* **[[Java对比|Java 对比]]**: Java 的 `ThreadLocal` 是隐式绑定的，容易导致内存泄漏；Go 的 `context` 必须作为函数的第一个参数显式传递。

### WaitGroup
用于等待一组并发操作完成。
* `Add(n)`: 计数加 n。
* `Done()`: 计数减 1。
* `Wait()`: 阻塞直到计数为 0。

### Mutex (互斥锁)
Go 的 `Mutex` 不支持重入（Reentrant）。
* **[[Java]] 对比**: Java 的 `synchronized` 和 `ReentrantLock` 允许同一个线程多次获取锁，但 Go 的 `Mutex` 如果在同一个 goroutine 中重复 Lock 会导致死锁。

## 4. 代码示例
请参考示例文件：[examples/main.go](examples/main.go)

```go
// 典型的 context 传值用法
ctx := context.WithValue(context.Background(), "userID", 123)
val := ctx.Value("userID")
```

## 5. 常见误区
* **Context 存储可选参数**: 不要把业务参数（如用户名、密码）放在 context 里，它应该只存放请求范围的数据（如 TraceID、用户身份信息）。
* **忘记调用 cancel()**: `WithTimeout` 等函数返回的 `cancel` 函数必须被调用，否则会导致 context 泄漏。
* **Mutex 重入**: 再次强调，Go 的 `Mutex` 不是可重入锁。
* **WaitGroup 计数为负**: 如果调用 `Done()` 的次数超过了 `Add()` 的次数，程序会 [[panic]]。

## 6. 本节练习
1. 编写一个 HTTP 客户端请求，使用 `context.WithTimeout` 设置 500ms 超时。
2. 使用 `sync.RWMutex` 实现一个简单的 [[并发安全]]缓存（[[Map]]）。
3. 使用 `sync.Once` 实现一个单例模式。

## 7. 面试可能怎么问
* **问**: 为什么 Go 推荐显式传递 context 而不是像 Java 那样使用 [[ThreadLocal]]？
* **答**: 显式传递使代码逻辑更清晰，开发者能明确知道哪些函数受 context 控制。同时避免了 ThreadLocal 在异步环境下难以追踪和清理的问题。
* **问**: 如何停止一个正在运行的 goroutine？
* **答**: 通过监听 `context.Done()` 信号或使用专门的退出 [[Channel|channel]]。Go 不支持从外部强制杀掉一个 goroutine。
* **问**: `sync.Mutex` 和 `sync.RWMutex` 有什么区别？
* **答**: `Mutex` 是完全互斥的；`RWMutex` 允许多个读锁并发，但写锁是互斥的，适用于读多写少的场景。

## 8. 本节总结
`context` 是 Go 处理并发请求的标准方式。配合 `sync` 包提供的同步原语，你可以构建出既高效又健壮的并发系统。记住：优先使用 channel 通信，只有在处理共享状态或简单的同步需求时才考虑使用锁。
