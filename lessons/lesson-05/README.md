---
title: "Lesson 05: struct / method / interface"
chapter: 05
part: Go语言基础
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 对象模型
prerequisites:
  - "[[lesson-04/README]]"
related:
  - "[[lesson-04/README]]"
  - "[[lesson-06/README]]"
key_concepts:
  - struct
  - method
  - interface
  - receiver
  - 组合
  - 多态
---

# Lesson 05: [[struct]] / [[method]] / [[interface]]

## 1. 学习目标
- 理解 [[Go语言|Go 语言]]中 struct 与 [[Java]] class 的本质区别
- 掌握 method 的定义方式及 [[receiver]] 的选择
- 深刻理解 Go 的隐式接口实现机制
- 掌握组合（Composition）在 Go 中的应用

## 2. 给 [[Java开发者|Java 开发者]]的类比

### struct vs class
Java 的 `class` 既包含数据也包含行为。Go 的 `struct` 只包含数据。行为是通过在 `struct` 之外定义 `method` 并绑定到它上面的。

### 隐式接口 vs 显式接口
Java 需要 `implements InterfaceName`。Go 不需要。只要你的类型实现了接口定义的所有方法，你就自动实现了该接口。这被称为“鸭子类型”（Duck Typing）。

### 组合 vs 继承
Java 使用 `extends`。Go 使用“匿名嵌入”。如果你在 `struct A` 中嵌入了 `struct B`，那么 `A` 自动拥有了 `B` 的所有字段和方法。

## 3. 核心概念

### Receiver (接收者)
- **Value Receiver**: 方法操作的是对象的副本，不会修改原对象。
- **Pointer Receiver**: 方法操作的是对象的指针，可以修改原对象，且避免了大对象的拷贝开销。

### 接口的本质
接口在 Go 中是一组方法的集合。它是非侵入式的，这使得你可以为第三方库的类型定义接口，而无需修改库的源码。

### 空接口 any
`interface{}`（现在常用 `any` 关键字）可以代表任何类型，类似于 Java 的 `Object`。但在使用时通常需要配合类型断言（Type Assertion）。

### [[Java对比|Java 对比]]
- **多态**: Java 的多态是基于类继承体系的。Go 的多态是基于接口满足的，更加解耦。
- **构造函数**: Go 没有构造函数，通常使用 `NewXxx` 命名的普通函数来返回初始化的结构体指针。

## 4. 代码示例
参考示例文件：`examples/interfaces.go`

```go
type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

// 指针接收者实现接口
func (c *Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}
```

## 5. 常见误区
- **误区 1**: 试图在 struct 中写方法。方法必须写在 struct 定义之外。
- **误区 2**: 混淆值接收者和指针接收者。如果你需要修改状态，必须用指针。
- **误区 3**: 过度使用 `any`。这会丢失类型安全性，增加运行时的断言开销。

## 6. 本节练习
1. 定义一个 `Animal` 接口，包含 `Eat()` 方法。让 `Cat` 和 `Dog` 结构体分别实现它。
2. 编写一个函数，接收 `any` 类型的参数，并根据其实际类型打印不同的信息。
3. 尝试使用结构体嵌套实现一个简单的“继承”效果，并调用被嵌入结构体的方法。

## 7. 面试可能怎么问
- **问**: Go 接口的隐式实现有什么好处？
- **答**: 它极大地降低了代码的耦合度。你可以在不修改原有代码的情况下，为现有的类型增加新的接口抽象。这对于大型项目的重构和模块化非常有利。
- **问**: 什么时候该用指针接收者？
- **答**: 1. 需要修改接收者的状态时；2. 接收者结构体很大，为了避免值拷贝的性能开销时；3. 为了保持一致性，如果该类型其他方法用了指针接收者，建议全部统一。

## 8. 本节总结
Go 通过 struct 和 interface 的组合，提供了一种比 Java 继承体系更灵活、更简洁的对象模型。记住：不要去模拟类，要去思考行为的抽象。
