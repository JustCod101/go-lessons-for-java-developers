---
title: "Lesson 12: HTTP 服务 (HTTP Server)"
chapter: 12
part: Go高级特性
date: 2026-05-04
status: published
tags:
  - Go
  - Java对比
  - Web开发
prerequisites:
  - "[[lesson-11/README]]"
related:
  - "[[lesson-11/README]]"
  - "[[lesson-13/README]]"
key_concepts:
  - net/http
  - Handler
  - ServeMux
  - Middleware
  - JSON
  - REST API
---

# Lesson 12: [[HTTP服务|HTTP 服务]] ([[HTTP Server]])

在 [[Java]] 中，我们习惯使用 Spring Boot 来构建 Web 应用。在 Go 中，标准库 `net/http` 已经非常强大，足以构建高性能的 [[REST API]]。

## 1. 学习目标
* 掌握 `net/http` 标准库的基本用法
* 理解 `Handler` 和 `ServeMux` 的概念
* 学会编写中间件 ([[Middleware]])
* 掌握 [[JSON]] 请求解析与响应返回
* 了解如何构建简单的分层架构

## 2. 给 [[Java开发者|Java 开发者]]的类比
* **`net/http`**: 相当于 Servlet API + 内置的 Tomcat/Jetty。
* **`ServeMux`**: 相当于 Spring MVC 的 `DispatcherServlet` 或路由映射。
* **`HandlerFunc`**: 相当于 Controller 中的一个 `@RequestMapping` 方法。
* **`Middleware`**: 相当于 Spring 的 `Filter` 或 `HandlerInterceptor`。
* **`json.Marshal/Unmarshal`**: 相当于 Jackson 或 Gson。

## 3. 核心概念

### [[Handler]] 与 [[ServeMux]]
Go 的 HTTP 处理核心是 `http.Handler` 接口。`ServeMux` 是一个 HTTP 请求多路复用器，它根据 URL 匹配对应的 Handler。
* **[[Java对比|Java 对比]]**: 在 Spring Boot 中，你通过注解定义路由。在 Go 中，你通常显式地将路径注册到 `ServeMux`。

### 中间件 (Middleware)
Go 的中间件本质上是一个接收 `http.Handler` 并返回 `http.Handler` 的函数。它利用了闭包和装饰器模式。
* **Java 对比**: 类似于 AOP 或 Filter。Go 的中间件非常直观，没有复杂的配置，就是一层层的函数包装。

### JSON 处理
Go 使用结构体标签 (Struct Tags) 来控制 JSON 的序列化与反序列化。
* **Java 对比**: 类似于 Jackson 的 `@JsonProperty`。

### 简单分层
虽然 Go 没有强制要求，但通常我们会按照 Handler (Web层) -> Service (业务层) -> Repository (数据层) 进行拆分。
* **Java 对比**: 结构非常相似，但 Go 倾向于通过构造函数进行显式的依赖注入，而不是使用 `@Autowired`。

## 4. 代码示例

查看 `src/` 目录下的完整示例。

```go
func UserHandler(w http.ResponseWriter, r *http.Request) {
    // 解析请求 (类似于从 HttpServletRequest 获取数据)
    // 处理业务逻辑
    // 返回 JSON (类似于 @ResponseBody)
    user := User{ID: 1, Name: "Tom"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

运行示例：
```bash
go run src/main.go src/handler.go
```

## 5. 常见误区
* **寻找 "Go 版 Spring Boot"**: 很多 Java 开发者第一反应是找一个像 Spring 一样全家桶式的框架（如 Gin, Echo, Beego）。虽然这些框架很好用，但建议先掌握 `net/http`，因为它是所有框架的基础。
* **错误处理**: 在 Handler 中不要忘记检查错误。如果解析 JSON 失败，必须显式调用 `http.Error` 并返回，否则程序会继续执行。
* **[[并发安全]]**: 每个 HTTP 请求都在独立的 [[Goroutine]] 中运行。如果你的 Handler 访问了全局变量，必须考虑加锁（类似于 Java 中的单例 Bean 访问成员变量）。

## 6. 本节练习
1. 修改 `UserHandler`，支持通过 URL 参数（如 `/user?id=123`）获取 ID。
2. 编写一个简单的中间件，在每个请求的响应头中添加 `X-Server-Name: GoServer`。
3. 实现一个 POST 接口，接收 JSON 格式的 User 对象并打印出来。

## 7. 面试可能怎么问
* **Q: Go 如何处理并发的 HTTP 请求？**
  * A: `net/http` 服务器为每个到来的连接创建一个新的 Goroutine 来处理请求。
* **Q: 什么是中间件？如何实现一个？**
  * A: 中间件是一个函数，它包装一个 `http.Handler` 并返回一个新的 `http.Handler`。它可以在处理请求前后执行逻辑。
* **Q: 如何在 Go 中返回 404 错误？**
  * A: 使用 `http.Error(w, "Not Found", http.StatusNotFound)`。

## 8. 本节总结
Go 的 Web 开发非常透明。没有隐藏的魔法，每一个路由、每一个中间件都是你亲手注册和包装的。这种“显式优于隐式”的风格虽然代码量稍多，但极大地降低了调试难度。
