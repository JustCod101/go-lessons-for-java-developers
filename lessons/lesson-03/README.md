---
title: "Lesson 03: 基础语法"
chapter: 03
part: Go语言基础
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 基础语法
prerequisites:
  - "[[lesson-02/README]]"
related:
  - "[[lesson-02/README]]"
  - "[[lesson-04/README]]"
key_concepts:
  - 变量声明
  - 零值
  - 指针
  - defer
  - 流程控制
---

# Lesson 03: 基础语法

## 1. 学习目标
- 掌握 Go 的变量声明与零值机制
- 理解 Go 的基本类型与字符串处理
- 熟练使用 if、for、switch 等流程控制
- 初步了解指针与 [[defer]] 的用法

## 2. 给 [[Java开发者|Java 开发者]]的类比

### 变量声明：类型后置
Java 是 `String name = "Gopher"`，Go 是 `var name string = "Gopher"`。Go 认为变量名比类型更重要，所以放在前面。

### 流程控制：没有括号
Java 的 `if (x > 0) { ... }` 在 Go 中变成了 `if x > 0 { ... }`。虽然去掉了括号，但大括号 `{}` 是强制要求的，即使只有一行代码。

### 循环：万能的 for
Java 有 `for`、`while`、`do-while`。Go 只有 `for`。
- `for i := 0; i < 10; i++` (普通 for)
- `for x < 10` (等价于 while)
- `for { ... }` (等价于 while(true))

## 3. 核心概念

### 变量与零值
Go 声明变量后如果不赋值，会自动初始化为该类型的“零值”：
- 数值类型：`0`
- 布尔类型：`false`
- 字符串：`""`
- 指针/接口/[[Slice|切片]]：`nil`

### 指针基础
Go 支持指针，但不支持指针运算。这让你能控制内存布局，又不会像 C 语言那样容易写出内存安全漏洞。

### defer 延迟执行
`defer` 语句会将函数调用推迟到当前函数返回前执行。这非常适合用于释放资源（如关闭文件、解锁），类似于 [[Java]] 的 `finally` 块，但写起来更优雅。

### [[Java对比|Java 对比]]
- **基本类型**: Go 的 `int` 大小取决于平台（32位或64位），而 Java 的 `int` 永远是 32 位。
- **字符串**: Go 的字符串是不可变的字节序列，默认使用 UTF-8 编码。

## 4. 代码示例
参考示例文件：`examples/syntax.go`

```go
package main

import "fmt"

func main() {
    // 简短声明，只能在函数内部使用
    count := 10
    
    if count > 5 {
        fmt.Println("Greater than 5")
    }

    // defer 示例
    defer fmt.Println("This prints last")
    fmt.Println("This prints first")
}
```

## 5. 常见误区
- **误区 1**: 尝试使用 `while` 关键字。记住，Go 只有 `for`。
- **误区 2**: 在函数外部使用 `:=` 简短声明。外部必须使用 `var` 关键字。
- **误区 3**: 认为 `nil` 指针可以安全访问。访问 `nil` 指针会导致 [[panic]]，这和 Java 的 `NullPointerException` 一样致命。

## 6. 本节练习
1. 编写一个程序，使用 `for` 循环打印 1 到 100 之间的所有偶数。
2. 声明一个指针变量，修改它指向的值，并观察原变量的变化。
3. 尝试在一个函数中使用多个 `defer`，观察它们的执行顺序（提示：栈结构）。

## 7. 面试可能怎么问
- **问**: [[Go语言|Go 语言]]中的 `iota` 是做什么用的？
- **答**: `iota` 是常量计数器，只能在常量的表达式中使用。它在每个 `const` 关键字出现时被重置为 0，每出现一个常量，`iota` 就会自动加 1。常用于定义枚举。
- **问**: 为什么 Go 强制要求大括号不能换行？
- **答**: 这是为了保持代码风格的高度统一，同时也与 Go 编译器的分号自动插入机制有关。

## 8. 本节总结
Go 的语法非常克制。你会发现没有了繁琐的括号和冗余的关键字，代码变得更加清爽。指针的引入给了你操作内存的能力，而 `defer` 则让你不再担心资源泄露。
