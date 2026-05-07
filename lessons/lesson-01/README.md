---
title: "Lesson 01: Go 语言整体认知"
chapter: 01
part: Go语言基础
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 课程导论
prerequisites:
  - "[[课程介绍]]"
related:
  - "[[lesson-02/README]]"
key_concepts:
  - Go语言
  - Java对比
  - Goroutine
  - Channel
  - 组合
  - 显式错误处理
---

# Lesson 01: [[Go语言|Go 语言]]整体认知

## 1. 学习目标
- 了解 Go 语言的设计哲学与核心特性
- 理解 Go 与 [[Java]] 在应用场景和定位上的差异
- 掌握 Go 语言的核心关键词
- 完成并运行第一个 Hello World 程序

## 2. 给 [[Java开发者|Java 开发者]]的类比

### 运行机制：JVM vs 原生二进制
Java 代码运行在 JVM 上，需要先编译成字节码，再由 JIT 编译执行。Go 则是直接编译成机器码，生成单个静态链接的二进制文件。

### 编程范式：继承 vs 组合
Java 是典型的面向对象语言，强调类继承和多态。Go 抛弃了继承，通过结构体组合和隐式接口实现多态，更加灵活。

### [[并发模型]]：线程 vs [[Goroutine]]
Java 的线程通常对应操作系统的内核线程，开销较大。Go 引入了轻量级的 Goroutine，单机支持百万级并发，切换成本极低。

## 3. 核心概念

### Go 是什么
Go 是由 Google 开发的静态强类型、编译型语言。它旨在解决大规模软件开发中的效率、并发和维护问题。

### 核心关键词
- **Simple (简洁)**: 语法极简，没有冗余的特性，全语言只有 25 个关键字。
- **Fast Compile (快速编译)**: 编译速度极快，大型项目也能在秒级完成。
- **Concurrency (并发)**: 原生支持并发，通过 Goroutine 和 [[Channel]] 实现。
- **Explicit Error (显式错误处理)**: 不使用 Exception，而是通过多返回值显式处理错误。
- **Composition (组合)**: 提倡组合优于继承。
- **Standard Library (标准库)**: 拥有强大且稳定的标准库，尤其是网络编程方面。

### [[Java对比|Java 对比]]
- **定位**: Java 适合复杂的企业级应用，Go 适合云原生、微服务和高性能后端。
- **内存管理**: 两者都有 [[GC]]，但 Go 的 GC 停顿时间通常更短，且支持指针运算（受限）。

## 4. 代码示例
参考示例文件：`examples/hello.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go! 欢迎来到 Go 语言的世界。")
}
```

## 5. 常见误区
- **误区 1**: 认为 Go 是 C++ 的替代品。实际上 Go 的开发效率更接近 Python，而性能接近 C/C++。
- **误区 2**: 寻找 Go 中的 `class`。Go 没有类，只有 `struct`。
- **误区 3**: 期待 `try-catch`。Go 强制你检查每一个可能出错的函数返回值。

## 6. 本节练习
1. 修改 `hello.go`，尝试打印你的名字。
2. 使用 `go run hello.go` 运行程序。
3. 使用 `go build hello.go` 编译程序，观察生成的二进制文件大小。

## 7. 面试可能怎么问
- **问**: 为什么 Go 语言没有继承？
- **答**: Go 团队认为继承增加了代码的耦合度和复杂性。通过组合和接口，可以实现更清晰、更灵活的代码结构。
- **问**: Go 的并发为什么比 Java 强？
- **答**: [[Thread|Java 线程]]是内核级的，创建和切换开销大。Go 的 Goroutine 是用户态线程，初始栈仅 2KB，由 Go 运行时调度，效率极高。

## 8. 本节总结
Go 语言不是为了取代 Java，而是为了在云原生时代提供更高效的开发体验。它通过减法设计，去掉了 Java 中许多复杂的特性，让我们能更专注于业务逻辑本身。
