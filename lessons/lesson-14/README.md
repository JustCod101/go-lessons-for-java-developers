---
title: "Lesson 14: 综合项目 Task API (Mini Project)"
chapter: 14
part: Go高级特性
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - 综合项目
prerequisites:
  - "[[lesson-13/README]]"
related:
  - "[[lesson-13/README]]"
  - "[[lesson-15/README]]"
key_concepts:
  - RESTful API
  - HTTP Handler
  - 并发安全
  - Repository
  - httptest
  - 集成测试
---

# Lesson 14: 综合项目 Task API (Mini Project)

这是本课程的最后一个阶段。我们将结合前面学到的 [[HTTP服务|HTTP 服务]]、[[JSON]] 处理、[[并发安全]]和测试知识，构建一个简单的任务管理 (To-Do List) API。

## 1. 学习目标
* 综合运用 `net/http` 构建 [[RESTful API]]
* 掌握如何在 Go 中组织项目结构
* 理解并发安全的内存存储实现
* 学会编写集成测试 (Integration Test)
* 体验从零开始构建一个完整的小型服务

## 2. 给 [[Java开发者|Java 开发者]]的类比
* **项目结构**: 类似于一个精简版的 Spring Boot 项目，但没有 `pom.xml` 的繁琐配置。
* **`httptest`**: 相当于 `MockMvc`，用于模拟 HTTP 请求并验证响应。
* **`sync.RWMutex`**: 相当于 `ReentrantReadWriteLock`，用于保护共享资源。

## 3. 核心概念

### 项目结构 (Project Structure)
在 Go 中，我们通常按功能或层级组织代码。本项目采用了简单的扁平结构，但在大型项目中，你可能会看到 `cmd/` (入口), `internal/` (私有逻辑), `pkg/` (公共库) 等目录。
* **[[Java对比|Java 对比]]**: 类似于 Java 的包结构，但 Go 的包名通常与目录名一致。

### 内存存储与并发
由于 Go 的 [[HTTP Handler]] 是并发运行的，访问共享的 `map` 必须加锁。
* **[[Java]] 对比**: 类似于在单例 Service 中使用 `ConcurrentHashMap` 或手动加锁。

### RESTful 路由实现
在不使用框架的情况下，我们需要手动解析 URL 路径来分发请求。
* **Java 对比**: 类似于手写一个 `HttpServlet` 并根据 `request.getPathInfo()` 进行 `switch-case`。

### 集成测试
使用 `net/http/httptest` 包，我们可以不启动真实的服务器就测试整个 Handler 链路。
* **Java 对比**: 类似于使用 `@SpringBootTest` 配合 `TestRestTemplate` 或 `MockMvc`。

## 4. 代码示例

查看 `src/` 目录下的完整实现。

**核心逻辑：**
1. `model.go`: 定义 `Task` 结构体。
2. `repository.go`: 实现并发安全的内存存储。
3. `handler.go`: 处理 HTTP 请求路由与业务逻辑。
4. `main.go`: 组装依赖并启动服务。

运行项目：
```bash
go run src/*.go
```

测试项目：
```bash
go test -v src/*.go
```

## 5. 常见误区
* **全局变量滥用**: 尽量通过结构体字段传递依赖（如 `Handler` 持有 `Repository`），而不是直接访问全局变量。这有利于单元测试和代码解耦。
* **忽略 URL 参数校验**: 在手动解析 ID 时，务必处理 `strconv.Atoi` 可能返回的错误。
* **忘记设置 Content-Type**: 返回 JSON 时，记得设置 `w.Header().Set("Content-Type", "application/json")`。

## 6. 本节练习
1. 为 Task 增加一个 `Priority` 字段（整数 1-3）。
2. 实现“查询所有已完成任务”的过滤功能。
3. 尝试将内存存储替换为 Lesson 13 学过的 SQLite 存储。

## 7. 面试可能怎么问
* **Q: 为什么在 Repository 中要使用 `sync.RWMutex` 而不是普通的 `sync.Mutex`？**
  * A: `RWMutex` 支持多读单写。对于任务列表这种读多写少的场景，性能更好。
* **Q: 如何测试一个 Go 的 HTTP Handler？**
  * A: 使用 `httptest.NewRequest` 创建请求，使用 `httptest.NewRecorder` 记录响应，然后调用 Handler 的 `ServeHTTP` 方法。
* **Q: [[Go项目|Go 项目]]中 `internal` 目录的作用是什么？**
  * A: `internal` 目录下的代码只能被其父目录及其子目录的代码导入，其他项目无法引用。这用于封装私有逻辑。

## 8. 本节总结
恭喜你完成了 [[Go语言|Go 语言]]的入门课程！通过这个小项目，你应该已经感受到了 Go 的简洁与高效。虽然它没有 Java 那么多开箱即用的魔法，但它赋予了开发者更清晰的视野和更强的掌控力。继续探索，你会发现 Go 在云原生和微服务领域的巨大魅力。
