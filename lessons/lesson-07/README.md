---
title: "Lesson 07: 包与项目结构"
chapter: 07
part: Go语言基础
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 项目结构
prerequisites:
  - "[[lesson-06/README]]"
related:
  - "[[lesson-06/README]]"
  - "[[lesson-08/README]]"
key_concepts:
  - Go Module
  - package
  - 可见性
  - internal
  - cmd
  - pkg
---

# LESSON 07: 包与项目结构

在 [[Java]] 中，我们习惯了深层的 [[package]] 嵌套（如 `com.company.project.service.impl`）。Go 的哲学完全不同，它提倡扁平化和简洁性。

## 1. 学习目标
* 理解 [[Go Module]] 的基本概念。
* 掌握包（package）的声明与导入规则。
* 理解首字母大小写控制可见性的机制。
* 熟悉 [[Go项目|Go 项目]]的常见目录结构（cmd, internal, pkg）。
* 学习如何将 Java 的多层架构思想迁移到 Go。

## 2. 给 [[Java开发者|Java 开发者]]的类比
* **Package**: 类似于 Java 的 package，但 Go 的包名通常与目录名一致。
* **Module**: 类似于 [[Maven]] 的 `pom.xml` 或 [[Gradle]] 的 `build.gradle`，定义了项目的根路径和依赖。
* **Visibility**: Go 没有 `public/private/protected` 关键字。首字母大写即为 public，首字母小写即为 private。
* **Internal**: 类似于 Java 9+ 的模块化系统，限制某些包只能在本项目内使用。

## 3. 核心概念

### [[Go Modules]]
`go.mod` 文件是项目的核心。它定义了模块路径，这是导入本项目内其他包的基础。
* **[[Java对比|Java 对比]]**: 类似于 Maven 的 `groupId` 和 `artifactId`。

### 可见性 (Visibility)
这是 Go 最独特的特性之一。
* `func DoSomething()`: 外部包可见。
* `func doSomething()`: 仅当前包可见。
* **Java 对比**: 这种设计极大地减少了样板代码，不再需要写大量的 `public` 关键字。

### 目录结构惯例
* **cmd/**: 存放程序的入口（main 函数）。每个子目录对应一个可执行文件。
* **internal/**: 存放不希望被外部项目引用的代码。Go 编译器会强制执行这一规则。
* **pkg/**: 存放可以被外部项目引用的库代码（存在争议，有些开发者倾向于直接放在根目录）。

## 4. 代码示例
请参考示例项目：[examples/](examples/)

```go
// 导入本项目内的包
import "example.com/lesson07/internal/auth"
```

## 5. 常见误区
* **过度分层**: Java 开发者容易把 `controller/service/repository` 那一套生搬硬套到 Go。在 Go 中，过多的层级会导致严重的循环依赖问题。
* **循环依赖**: Go 不允许包之间的循环引用。如果 A 引用 B，B 就不能引用 A。这通常意味着你的包划分不够合理。
* **包名冲突**: 避免给包起名为 `util` 或 `common` 这种模糊的名字。尽量使用具体的业务名称。

## 6. 本节练习
1. 初始化一个新的 Go 模块，并创建一个名为 `mathutils` 的包，包含一个导出的 `Add` 函数。
2. 尝试在 `internal` 目录外引用 `internal` 目录下的包，观察编译器的报错信息。
3. 重构一个简单的 Java Service 类，尝试用 Go 的扁平化结构重新组织它。

## 7. 面试可能怎么问
* **问**: Go 是如何控制成员可见性的？
* **答**: 通过首字母大小写。大写字母开头的标识符是导出的，小写字母开头的是未导出的。
* **问**: `internal` 目录的作用是什么？
* **答**: 它是一个特殊的目录名。Go 编译器会限制 `internal` 目录下的包只能被其父目录下的包导入，防止内部实现细节泄露给外部用户。
* **问**: 如何解决 Go 中的循环依赖？
* **答**: 提取公共接口、重构包结构或使用依赖注入。

## 8. 本节总结
Go 的项目结构强调“少即是多”。不要试图在项目开始时就构建复杂的层级。从扁平的结构开始，随着业务增长再进行合理的拆分。记住，清晰的边界比深层的抽象更重要。
