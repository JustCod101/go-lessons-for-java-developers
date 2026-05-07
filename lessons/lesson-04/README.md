---
title: "Lesson 04: 函数与错误处理"
chapter: 04
part: Go语言基础
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 错误处理
prerequisites:
  - "[[lesson-03/README]]"
related:
  - "[[lesson-03/README]]"
  - "[[lesson-05/README]]"
key_concepts:
  - 多返回值
  - 闭包
  - error
  - panic
  - recover
  - error wrapping
---

# Lesson 04: 函数与错误处理

## 1. 学习目标
- 掌握 Go 函数的多返回值与命名返回值特性
- 理解匿名函数与闭包的应用
- 深刻理解 Go 的错误处理哲学：Error is Value
- 掌握 error wrapping 与解包技巧

## 2. 给 [[Java开发者|Java 开发者]]的类比

### 多返回值 vs 包装类
Java 函数只能返回一个值。如果想返回多个，通常要定义一个 `Result` 类。Go 原生支持多返回值，最常见的模式是 `(result, error)`。

### Error vs Exception
Java 使用 `try-catch-finally` 处理异常。Go 认为错误是正常的业务逻辑分支，不是“异常”。你必须显式检查 `if err != nil`。

### 闭包 vs 匿名内部类
Java 的匿名内部类访问外部变量有 `final` 限制。Go 的闭包更加自然，可以直接捕获并修改外部变量。

## 3. 核心概念

### 函数特性
- **多返回值**: 函数可以返回多个不同类型的值。
- **命名返回值**: 可以在函数签名中定义返回变量名，函数内直接赋值后 `return` 即可。
- **一等公民**: 函数可以作为变量传递，也可以作为参数和返回值。

### 错误处理哲学
Go 鼓励“尽早返回”。如果函数出错，立即返回错误，减少嵌套。
```go
f, err := os.Open("test.txt")
if err != nil {
    return err
}
// 正常逻辑继续...
```

### Error Wrapping
使用 `fmt.Errorf("...: %w", err)` 可以将原始错误包装起来，保留上下文信息，同时允许调用者使用 `errors.Is` 或 `errors.As` 进行判断。

### [[Java对比|Java 对比]]
- **性能**: `panic/recover` 机制类似于 Exception，但开销巨大。Go 推荐只在不可恢复的灾难性错误（如[[Array|数组]]越界）时使用 `panic`。
- **代码量**: 虽然 `if err != nil` 增加了代码行数，但它让程序的执行路径变得极其清晰。

## 4. 代码示例
参考示例文件：`examples/errors.go`

```go
func process() (int, error) {
    result, err := doSomething()
    if err != nil {
        return 0, fmt.Errorf("process failed: %w", err)
    }
    return result, nil
}
```

## 5. 常见误区
- **误区 1**: 滥用 `panic`。在 Go 中，`panic` 应该极少出现。
- **误区 2**: 忽略错误。永远不要写 `_ = doSomething()`，除非你百分之百确定错误不重要。
- **误区 3**: 忘记处理命名返回值的覆盖问题。

## 6. 本节练习
1. 编写一个函数，接收两个整数并返回它们的商和余数。
2. 实现一个简单的闭包计数器。
3. 模拟一个文件读取过程，如果文件不存在，返回一个包装后的自定义错误。

## 7. 面试可能怎么问
- **问**: 为什么 Go 不支持 try-catch？
- **答**: Go 的设计者认为 try-catch 容易导致开发者忽略错误，且会使控制流变得复杂。显式错误处理强制开发者思考每一个失败场景，提高代码健壮性。
- **问**: `errors.Is` 和 `errors.As` 有什么区别？
- **答**: `errors.Is` 用于判断错误是否为特定实例（类似 `==`）；`errors.As` 用于判断错误是否为特定类型并提取其值（类似类型转换）。

## 8. 本节总结
函数是 Go 的基本构建块。理解了“错误即值”的理念，你就掌握了 [[Go编程|Go 编程]]的灵魂。不要抱怨 `if err != nil` 繁琐，它是你程序稳定运行的守护神。
