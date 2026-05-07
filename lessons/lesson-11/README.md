---
title: "Lesson 11: 测试 (Testing)"
chapter: 11
part: Go高级特性
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 测试
prerequisites:
  - "[[lesson-10/README]]"
related:
  - "[[lesson-10/README]]"
  - "[[lesson-12/README]]"
key_concepts:
  - testing
  - 单元测试
  - 表驱动测试
  - Benchmark
  - Mock
  - JUnit
---

# Lesson 11: 测试 (Testing)

在 [[Java]] 世界里，[[JUnit]] 是测试的标准。在 Go 中，测试是语言内置的一等公民，不需要引入第三方库就能完成大部分工作。

## 1. 学习目标
* 掌握 `testing` 包的基本用法
* 学会编写单元测试和表驱动测试 (Table-Driven Tests)
* 了解如何进行性能测试 ([[Benchmark]])
* 掌握查看测试覆盖率的方法
* 理解 [[Go测试|Go 测试]]与 JUnit 的核心差异

## 2. 给 [[Java开发者|Java 开发者]]的类比
* **`testing` 包**: 相当于 JUnit 5 核心库。
* **`go test` 命令**: 相当于 `mvn test` 或 `gradle test`。
* **`TestXxx(t *testing.T)`**: 相当于带有 `@Test` 注解的方法。
* **`t.Errorf`**: 相当于 `Assertions.assertEquals`，但 Go 倾向于手动检查并报告错误。
* **`BenchmarkXxx(b *testing.B)`**: 相当于 JMH (Java Microbenchmark Harness)。

## 3. 核心概念

### 测试文件命名规则
Go 要求测试文件必须以 `_test.go` 结尾。如果你的代码在 `calc.go` 中，测试代码通常放在 `calc_test.go`。
* **[[Java对比|Java 对比]]**: Java 通常将测试放在 `src/test/java` 目录下，而 Go 习惯将测试文件与源代码放在同一个包（目录）下。

### 单元测试 (Unit Test)
测试函数签名必须是 `func TestXxx(t *testing.T)`。
* **Java 对比**: Go 没有 `@BeforeEach` 或 `@AfterEach` 这样的注解。如果需要 Setup/Teardown，通常在函数内部手动处理，或者使用 `TestMain`。

### 表驱动测试 (Table-Driven Tests)
这是 Go 的特色。通过定义一个结构体[[Slice|切片]]（[[Array|数组]]），循环运行测试用例。
* **Java 对比**: 类似于 JUnit 5 的 `@ParameterizedTest`，但 Go 的实现更显式、更灵活，不需要复杂的注解配置。

### 性能测试 (Benchmark)
函数签名必须是 `func BenchmarkXxx(b *testing.B)`。
* **Java 对比**: Java 需要引入 JMH 这种复杂的框架，而 Go 直接内置。

### [[Mock]] 的基本思路
Go 没有像 Mockito 那样强大的字节码增强库。Go 的 Mock 主要依赖于 **接口 ([[interface|Interface]])**。
* **Java 对比**: 在 Java 中你可能习惯用 `any()` 或 `when().thenReturn()`。在 Go 中，你需要定义接口，然后在测试中传入一个实现了该接口的 Mock 结构体。

## 4. 代码示例

查看 `examples/calc_test.go` 获取完整代码。

```go
func TestAdd(t *testing.T) {
    got := Add(1, 2)
    want := 3
    if got != want {
        t.Errorf("Add(1, 2) = %d; want %d", got, want)
    }
}
```

运行测试：
```bash
go test -v .
```

查看覆盖率：
```bash
go test -cover .
```

## 5. 常见误区
* **过度依赖 Mock 框架**: 刚从 Java 转过来的开发者总想找 "Go 版的 Mockito"。其实在 Go 中，简单的接口替换通常比引入复杂的 Mock 框架更清晰。
* **断言库的选择**: Go 标准库没有 `assert` 函数。虽然有 `testify/assert` 这样的第三方库，但官方建议先学会使用原生的 `if` 判断和 `t.Errorf`。
* **测试私有函数**: 在 Go 中，测试文件如果和源代码在同一个 [[package]]，可以直接测试私有（小写开头）函数。

## 6. 本节练习
1. 为 `Add` 函数增加更多的边界情况测试（如最大整数）。
2. 编写一个 `Subtract` 函数并为其编写表驱动测试。
3. 尝试运行 `go test -bench=.` 看看 Benchmark 的输出。

## 7. 面试可能怎么问
* **Q: Go 的测试文件命名有什么要求？**
  * A: 必须以 `_test.go` 结尾。
* **Q: 什么是表驱动测试？为什么要用它？**
  * A: 通过结构体数组定义输入和预期输出，在循环中调用 `t.Run`。它能减少重复代码，使测试用例更易于维护和扩展。
* **Q: 如何在 Go 中做 Mock？**
  * A: 主要通过接口。定义依赖为接口类型，在测试时传入一个实现了该接口的 Mock 对象。

## 8. 本节总结
Go 的测试哲学是“简单、显式”。它不依赖魔法般的注解，而是通过普通的 Go 代码来完成测试逻辑。掌握了表驱动测试，你就掌握了 Go 测试的精髓。
