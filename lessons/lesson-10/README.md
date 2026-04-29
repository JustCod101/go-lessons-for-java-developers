# LESSON 10: 标准库

Go 语言被称为“自带电池”（Batteries Included）的语言。它的标准库非常强大且设计精良，许多在 Java 中需要第三方库（如 Apache Commons, Jackson, Spring）的功能，在 Go 中直接使用标准库就能搞定。

## 1. 学习目标
* 熟悉 Go 标准库的常用包及其功能。
* 掌握 `fmt` 进行格式化输出。
* 学习 `time` 包处理时间与定时任务。
* 掌握 `encoding/json` 进行序列化与反序列化。
* 学习使用 `net/http` 构建基础的 Web 服务。
* 了解 `os` 和 `io` 包进行文件与系统交互。

## 2. 给 Java 开发者的类比
* **fmt**: 类似于 `System.out.printf` 或 `String.format`。
* **encoding/json**: 类似于 `Jackson` 或 `Gson`。
* **net/http**: 类似于内置了 `Tomcat/Jetty` 的 `Spring Boot`，但更轻量、更底层。
* **time**: 类似于 `java.time` 包。
* **os/io**: 类似于 `java.io` 和 `java.nio`。

## 3. 核心概念

### 格式化 (fmt)
Go 的 `fmt` 包使用了类似于 C 语言的格式化占位符。
* `%v`: 默认格式输出。
* `%+v`: 结构体输出时包含字段名。
* `%T`: 输出变量类型。

### JSON 处理 (encoding/json)
Go 通过结构体标签（Struct Tags）来控制 JSON 的映射关系。
* **Java 对比**: 类似于 Jackson 的 `@JsonProperty` 注解。

### HTTP 服务 (net/http)
Go 原生支持高性能的 HTTP 服务。你不需要安装额外的 Web 容器。
* **Java 对比**: 在 Java 中你可能需要引入 Spring MVC 和 Tomcat；在 Go 中，几行代码就能启动一个生产级别的 HTTP Server。

### 时间处理 (time)
Go 的时间格式化非常独特，它不使用 `yyyy-MM-dd`，而是使用一个固定的参考时间：`2006-01-02 15:04:05`。
* **记忆窍门**: 1 (月) 2 (日) 3 (时) 4 (分) 5 (秒) 6 (年)。

## 4. 代码示例
请参考示例文件：[examples/main.go](examples/main.go)

```go
// 快速解析 JSON
var user User
err := json.Unmarshal([]byte(jsonStr), &user)
```

## 5. 常见误区
* **时间格式化字符串**: 必须使用 `2006-01-02 15:04:05` 这个特定的时间点，否则格式化结果会完全错误。
* **忽略错误处理**: 标准库中的大多数函数都会返回 `error`，Java 开发者习惯了 Exception，容易忘记检查这些错误。
* **HTTP Body 未关闭**: 在发送 HTTP 请求后，必须手动调用 `resp.Body.Close()`，否则会导致连接泄漏。

## 6. 本节练习
1. 使用 `os` 包读取一个本地文件，并统计其中的单词数量。
2. 编写一个简单的 HTTP Server，接收 JSON 请求并返回处理后的结果。
3. 使用 `time.Tick` 实现一个每隔 2 秒打印一次当前时间的定时器。

## 7. 面试可能怎么问
* **问**: Go 的 JSON 解析是如何处理私有字段的？
* **答**: `encoding/json` 包只能访问导出的字段（首字母大写）。如果字段名是小写的，它将被忽略。
* **问**: 如何在 Go 中实现一个简单的 HTTP 中间件？
* **答**: 中间件本质上是一个接收 `http.Handler` 并返回 `http.Handler` 的函数。
* **问**: Go 的时间格式化为什么设计成 2006-01-02 这种形式？
* **答**: 这种设计避免了 `yyyy` 这种占位符的歧义，通过具体的数字示例来定义格式，更加直观（虽然初学者需要时间适应）。

## 8. 本节总结
Go 的标准库是学习 Go 最佳实践的宝库。它不仅功能强大，而且代码风格高度统一。深入理解标准库，不仅能提高开发效率，还能让你写出更地道的 Go 代码。
