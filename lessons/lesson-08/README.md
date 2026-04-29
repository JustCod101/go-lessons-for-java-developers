# LESSON 08: goroutine 和 channel

Go 语言最引以为傲的特性就是其原生支持的并发模型。不同于 Java 的共享内存模型，Go 提倡“通过通信来共享内存，而不是通过共享内存来通信”。

## 1. 学习目标
* 理解 goroutine 的轻量级特性及其调度原理。
* 掌握 channel 的基本操作（发送、接收、关闭）。
* 区分有缓冲（buffered）与无缓冲（unbuffered）channel。
* 熟练使用 `select` 语句处理多路复用。
* 学习常见的并发模式（Worker Pool, Fan-in/Fan-out）。
* 识别并避免死锁与 goroutine 泄漏。

## 2. 给 Java 开发者的类比
* **goroutine**: 类似于 Java 21 引入的虚拟线程（Virtual Threads/Project Loom）。它们都非常轻量，可以在一个 OS 线程上运行成千上万个。
* **channel**: 类似于 `BlockingQueue`。它不仅是数据传输的通道，还起到了同步的作用。
* **select**: 类似于 Java NIO 中的 `Selector`，但它用于监听 channel 的操作，语法更简洁。

## 3. 核心概念

### Goroutine
使用 `go` 关键字即可启动一个 goroutine。它的初始栈空间仅为 2KB 左右，而 Java 线程通常需要 1MB。
* **Java 对比**: Java 线程是重量级的，通常需要线程池来管理。Go 虽然也有协程池的概念，但大多数情况下直接 `go` 即可。

### Channel
Channel 是 goroutine 之间的通信桥梁。
* **无缓冲 Channel**: 发送和接收必须同时准备好，否则会阻塞。这是一种强同步机制。
* **有缓冲 Channel**: 类似于固定容量的队列。只有当缓冲区满时，发送才会阻塞。

### Select
`select` 允许一个 goroutine 等待多个 channel 操作。如果多个 case 同时就绪，它会随机选择一个执行。

## 4. 代码示例
请参考示例文件：[examples/main.go](examples/main.go)

```go
// 简单的超时控制模式
select {
case res := <-c:
    fmt.Println(res)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

## 5. 常见误区
* **向已关闭的 channel 发送数据**: 会触发 panic。
* **从已关闭的 channel 接收数据**: 会立即返回该类型的零值。可以通过 `v, ok := <-ch` 中的 `ok` 来判断 channel 是否已关闭。
* **Goroutine 泄漏**: 如果一个 goroutine 在等待一个永远不会有数据的 channel，它将永远驻留在内存中。
* **死锁**: 所有的 goroutine 都在等待，没有一个在运行。最常见的是在主协程中进行同步读写而没有其他协程配合。

## 6. 本节练习
1. 实现一个简单的生产者-消费者模型。
2. 编写一个程序，启动 10 个 goroutine 并发计算，最后汇总结果。
3. 模拟一个 goroutine 泄漏的场景，并思考如何使用 `context`（下一节课内容）来修复它。

## 7. 面试可能怎么问
* **问**: goroutine 和线程有什么区别？
* **答**: goroutine 是用户态线程，由 Go 运行时调度；线程由操作系统调度。goroutine 栈空间动态增长，切换成本极低。
* **问**: 什么是无缓冲 channel 和有缓冲 channel？
* **答**: 无缓冲 channel 是同步的，发送和接收必须配对；有缓冲 channel 是异步的，在容量范围内不会阻塞。
* **问**: 如何优雅地关闭 channel？
* **答**: 遵循“由发送方关闭”的原则。如果有多个发送方，通常需要引入额外的信号 channel 或使用 `sync.WaitGroup`。

## 8. 本节总结
掌握 goroutine 和 channel 是从 Java 开发者转型为 Go 开发者的分水岭。不要试图用 Java 的 `synchronized` 或 `Lock` 来解决所有问题，尝试拥抱 channel，你会发现并发编程可以变得如此优雅和简单。
