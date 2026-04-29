````markdown
# OpenCode UltraWork Prompt：为有 Java 基础的 0 基础 Go 学习者编写系统化 Lessons

你现在进入 **UltraWork 模式**。

## 一、你的角色

你是一位资深 Go 语言工程师、Java 后端架构师、技术课程设计专家。

你非常熟悉：

- Java 与 Go 的语言范式差异
- Java 后端开发者学习 Go 时最容易踩的坑
- Go 的工程化开发方式
- Go Modules、并发模型、接口设计、错误处理、测试、标准库、Web 开发、数据库访问、微服务开发
- 如何从 0 到 1 设计可执行、可练习、可沉淀的编程课程

你的目标不是写一本泛泛而谈的 Go 教程，而是为「有 Java 基础，但 Go 零基础」的人设计一套 **高质量、项目驱动、可落地、适合自学与实战的 Go Lessons**。

---

## 二、目标用户画像

学习者具备：

- Java 基础语法
- 面向对象编程经验
- 了解集合、异常、泛型、接口、线程、JVM、Spring Boot 等基础概念
- 可能做过 Java 后端项目
- 对 Go 完全零基础

学习目标：

- 快速掌握 Go 语言核心语法
- 理解 Go 与 Java 的底层思想差异
- 能独立阅读 Go 项目源码
- 能使用 Go 编写 CLI、HTTP Server、REST API、数据库 CRUD、小型微服务
- 能理解 Go 的并发模型 goroutine / channel / context
- 能达到后端实习 / 初级 Go 工程师的基础水平

---

## 三、任务目标

请你在当前项目目录中创建一套完整的 Go 学习课程，命名为：

```text
go-lessons-for-java-developers
````

课程需要以 Markdown 文档 + Go 代码示例 + 练习题 + 小项目的形式组织。

最终输出应该是一套可以直接交给学习者自学的课程材料。

---

## 四、目录结构要求

请生成如下目录结构：

```text
go-lessons-for-java-developers/
├── README.md
├── roadmap.md
├── learning-method.md
├── java-vs-go.md
├── lessons/
│   ├── lesson-01-go-overview/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-02-environment-and-modules/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-03-basic-syntax/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-04-functions-and-error-handling/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-05-struct-method-interface/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-06-collections/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-07-packages-and-project-layout/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-08-goroutine-and-channel/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-09-context-and-concurrency-control/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-10-standard-library/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-11-testing/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-12-http-server/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-13-database-crud/
│   │   ├── README.md
│   │   └── examples/
│   ├── lesson-14-mini-project-task-api/
│   │   ├── README.md
│   │   └── src/
│   ├── lesson-15-mini-project-cache/
│   │   ├── README.md
│   │   └── src/
│   └── lesson-16-next-step/
│       └── README.md
├── exercises/
│   ├── README.md
│   ├── basic/
│   ├── intermediate/
│   └── advanced/
├── projects/
│   ├── cli-todo/
│   ├── http-task-api/
│   ├── memory-cache/
│   └── concurrent-crawler/
└── cheat-sheets/
    ├── go-syntax-cheatsheet.md
    ├── java-to-go-mapping.md
    ├── go-concurrency-cheatsheet.md
    └── common-pitfalls.md
```

---

## 五、课程设计原则

请严格遵守以下原则：

### 1. 面向 Java 开发者讲 Go

每个知识点都要尽量和 Java 对比，例如：

* `class` vs `struct`
* `implements` vs 隐式接口实现
* `try-catch` vs 显式 `error`
* `Thread` / `ExecutorService` vs `goroutine`
* `BlockingQueue` vs `channel`
* `Maven/Gradle` vs `Go Modules`
* `package private/public` vs Go 大小写导出规则
* `List/Map/Set` vs slice/map/自实现 set
* `Spring Boot Controller` vs `net/http`
* `JUnit` vs Go testing
* `Exception` vs error wrapping
* `synchronized/Lock` vs mutex/channel

### 2. 不要只讲语法，要讲设计哲学

必须解释 Go 为什么这么设计，例如：

* 为什么 Go 不强调传统 OOP
* 为什么 Go 没有类继承
* 为什么错误处理是显式返回值
* 为什么接口是隐式实现
* 为什么 goroutine 比线程轻量
* 为什么 Go 项目强调简单目录结构
* 为什么 Go 倡导组合优于继承
* 为什么 Go 很适合云原生、网络服务、基础设施开发

### 3. 每节 Lesson 必须包含固定结构

每个 `lessons/lesson-xx/README.md` 必须包含：

```text
# Lesson 标题

## 1. 学习目标

## 2. 给 Java 开发者的类比

## 3. 核心概念

## 4. 代码示例

## 5. 常见误区

## 6. 本节练习

## 7. 面试可能怎么问

## 8. 本节总结
```

### 4. 每节必须有可运行代码

每个 lesson 的 `examples/` 目录中必须至少包含 1 到 3 个 `.go` 文件。

代码必须：

* 可以通过 `go run` 运行
* 有清晰注释
* 体现该节核心知识点
* 不要过度复杂
* 不要依赖难以安装的第三方库，除非必要

### 5. 课程要循序渐进

课程要按照如下认知路径设计：

```text
语言概览
→ 环境与模块
→ 基础语法
→ 函数与错误处理
→ struct / method / interface
→ slice / map / collection
→ package 与项目结构
→ goroutine / channel
→ context / 并发控制
→ 标准库
→ 测试
→ HTTP 服务
→ 数据库 CRUD
→ 综合项目
→ 进阶路线
```

---

## 六、重点 Lessons 内容要求

### Lesson 01：Go 语言整体认知

需要讲清楚：

* Go 是什么
* Go 适合什么场景
* Go 不适合什么场景
* Go 与 Java 的定位差异
* Go 的核心关键词：

  * simple
  * fast compile
  * concurrency
  * explicit error
  * composition
  * standard library
* 第一个 Hello World

### Lesson 02：环境与 Go Modules

需要讲清楚：

* 安装 Go
* `go version`
* `go env`
* `go mod init`
* `go run`
* `go build`
* `go test`
* `go fmt`
* Go Modules 对比 Maven / Gradle
* `go.mod` 和 `go.sum`
* module、package、import 的关系

### Lesson 03：基础语法

需要覆盖：

* 变量声明
* 常量
* 基本类型
* 字符串
* if
* for
* switch
* defer 简介
* 指针基础
* Go 没有 while
* Go 的零值机制

### Lesson 04：函数与错误处理

需要覆盖：

* 函数定义
* 多返回值
* 命名返回值
* 匿名函数
* 闭包
* `error`
* `errors.New`
* `fmt.Errorf`
* error wrapping
* Java Exception 与 Go error 的根本区别
* 为什么 Go 不推荐把异常当控制流

### Lesson 05：struct / method / interface

需要重点讲清楚：

* Go 没有 class
* struct 是数据结构
* method 是绑定在类型上的函数
* receiver
* pointer receiver vs value receiver
* interface 是行为抽象
* 隐式实现接口
* 组合优于继承
* Java interface 和 Go interface 的差异
* 空接口 `any`

### Lesson 06：集合

需要覆盖：

* array
* slice
* slice 底层结构
* append
* len / cap
* map
* set 的自实现方式
* range
* slice 常见坑：

  * 共享底层数组
  * append 扩容
  * nil slice vs empty slice
* Java ArrayList / HashMap 对比

### Lesson 07：package 与项目结构

需要覆盖：

* package
* import
* module
* 大小写控制可见性
* internal 目录
* cmd 目录
* pkg 目录争议
* Go 项目常见结构
* Java 多层架构迁移到 Go 的方式
* Go 不推荐过度分层

### Lesson 08：goroutine 和 channel

需要重点讲清楚：

* goroutine 是什么
* goroutine vs Java Thread
* channel 是什么
* channel vs BlockingQueue
* buffered channel
* unbuffered channel
* select
* close channel
* worker pool
* fan-in / fan-out
* 常见死锁
* goroutine 泄漏

### Lesson 09：context 与并发控制

需要覆盖：

* context 是什么
* 为什么需要 context
* cancellation
* timeout
* deadline
* value
* context 在 HTTP 服务中的作用
* WaitGroup
* Mutex
* RWMutex
* Once
* atomic
* Java Future / CompletableFuture / ThreadLocal 对比

### Lesson 10：标准库

需要覆盖：

* fmt
* strings
* strconv
* time
* os
* io
* bufio
* encoding/json
* net/http
* log/slog
* flag
* regexp
* sort
* sync
* context

### Lesson 11：测试

需要覆盖：

* testing 包
* 单元测试
* 表驱动测试
* benchmark
* mock 的基本思路
* coverage
* Go 测试与 JUnit 的差异
* 测试文件命名规则
* 如何写可测试代码

### Lesson 12：HTTP Server

需要覆盖：

* `net/http`
* Handler
* ServeMux
* middleware
* JSON request / response
* REST API
* 参数解析
* 错误响应
* 简单分层：

  * handler
  * service
  * repository
* 对比 Spring Boot Controller

### Lesson 13：数据库 CRUD

需要覆盖：

* database/sql
* driver 概念
* MySQL 或 SQLite 示例
* 连接池
* Query / Exec
* Scan
* transaction
* repository pattern
* SQL 注入风险
* Java JDBC / MyBatis / JPA 对比

### Lesson 14：综合项目：Task API

实现一个小型任务管理 API：

功能：

* 创建任务
* 查询任务列表
* 根据 ID 查询任务
* 更新任务状态
* 删除任务
* JSON API
* 简单内存存储或 SQLite 存储
* 清晰项目结构
* 带测试

### Lesson 15：综合项目：内存缓存

实现一个简化版内存缓存：

功能：

* Set
* Get
* Delete
* TTL
* 过期清理
* 并发安全
* 使用 Mutex / RWMutex
* 简单 benchmark
* 对比 Java ConcurrentHashMap + ScheduledExecutorService

### Lesson 16：Go 进阶路线

需要给出后续学习路径：

* Go runtime
* GMP 调度模型
* GC
* pprof
* gRPC
* Gin / Echo / Fiber
* GORM / sqlc / ent
* Kubernetes client-go
* Docker / Kubernetes
* 微服务
* 分布式系统
* 云原生基础设施项目
* 推荐实战项目路线

---

## 七、练习题要求

在 `exercises/` 下生成分层练习：

### basic

至少 20 道基础练习，例如：

* 变量与类型
* if / for / switch
* 函数
* slice
* map
* struct
* method
* interface

### intermediate

至少 20 道中级练习，例如：

* 实现 Set
* 实现 Stack
* 实现 Queue
* JSON 编解码
* 文件读写
* HTTP handler
* 错误处理
* 表驱动测试
* goroutine 通信
* worker pool

### advanced

至少 15 道高级练习，例如：

* 并发爬虫
* 限流器
* TTL Cache
* 连接池模拟
* 简单 RPC
* graceful shutdown
* context 超时控制
* 日志中间件
* 简单任务调度器

每道题需要包含：

```text
题目
目标
输入输出示例
提示
参考解法路径
```

---

## 八、项目要求

在 `projects/` 下实现 4 个小项目：

### 1. cli-todo

命令行 Todo 工具：

* add
* list
* done
* delete
* JSON 文件持久化
* 使用 flag 包

### 2. http-task-api

HTTP REST API：

* CRUD
* JSON
* middleware
* error response
* graceful shutdown
* 单元测试

### 3. memory-cache

内存缓存：

* TTL
* 并发安全
* 定期清理
* benchmark
* README 说明设计权衡

### 4. concurrent-crawler

并发爬虫：

* worker pool
* channel
* context cancel
* timeout
* 去重
* 限制最大并发
* 输出抓取结果

---

## 九、Cheat Sheets 要求

生成 4 份速查表：

### go-syntax-cheatsheet.md

包含：

* 变量
* 常量
* 函数
* struct
* interface
* slice
* map
* error
* goroutine
* channel
* context
* testing

### java-to-go-mapping.md

必须包含 Java 到 Go 的对应关系表，例如：

```text
Java class -> Go struct + method
Java interface -> Go interface
Java exception -> Go error
Java ArrayList -> Go slice
Java HashMap -> Go map
Java Thread -> Go goroutine
Java ExecutorService -> goroutine + channel + WaitGroup
Java synchronized -> sync.Mutex
Java package private -> Go 小写标识符
Java public -> Go 大写标识符
Java Maven -> Go Modules
```

### go-concurrency-cheatsheet.md

包含：

* goroutine
* channel
* select
* WaitGroup
* Mutex
* RWMutex
* Once
* atomic
* context
* worker pool
* fan-in/fan-out
* timeout
* cancellation

### common-pitfalls.md

包含 Java 开发者学习 Go 的常见坑：

* 用 Java OOP 思维硬套 Go
* 过度抽象
* 到处定义 interface
* 忽视 error
* 滥用 panic
* goroutine 泄漏
* channel 不关闭或乱关闭
* slice 共享底层数组
* map 并发读写 panic
* 不理解 pointer receiver
* 不会写表驱动测试
* 不会使用 context
* 目录结构过度 Spring 化

---

## 十、代码质量要求

所有 Go 代码必须满足：

* 使用 Go 1.22+ 风格
* 代码可运行
* 命名清晰
* 注释适量
* 不写伪代码
* 不写无法编译的示例
* 优先使用标准库
* 必须包含必要的 `go.mod`
* 示例代码尽量短小，但必须完整
* 项目代码必须能通过：

  * `go fmt`
  * `go test`
  * `go run`

---

## 十一、README.md 要求

根目录 `README.md` 必须包含：

```text
# Go Lessons for Java Developers

## 这套课程适合谁

## 学完后你能获得什么能力

## 学习路线

## 如何使用本课程

## 推荐学习节奏

## 课程目录

## 项目实战

## Java 开发者学习 Go 的核心转变

## 常见问题
```

---

## 十二、roadmap.md 要求

请生成一份 4 周学习路线：

### 第 1 周

* Go 环境
* 基础语法
* 函数
* 错误处理
* struct

### 第 2 周

* interface
* collection
* package
* testing
* 标准库

### 第 3 周

* goroutine
* channel
* context
* HTTP Server
* 数据库 CRUD

### 第 4 周

* Task API
* Memory Cache
* Concurrent Crawler
* 总复盘
* 面试准备

每天需要给出：

* 学习内容
* 编码任务
* 复盘问题
* 预估耗时

---

## 十三、learning-method.md 要求

请专门写一份学习方法文档，告诉学习者如何用这套课程高效掌握 Go：

必须包含：

* 不要只看，要运行
* 每个 lesson 至少改 3 次代码
* 用 Java 类比，但不要被 Java 思维绑架
* 用小项目倒逼理解
* 每天写学习日志
* 每周做一次复盘
* 如何用 AI 辅助学习但不沦为“看过就忘”
* 如何通过输出倒逼输入
* 如何准备 Go 面试

---

## 十四、写作风格要求

整体风格：

* 中文
* 清晰
* 系统
* 对 Java 开发者友好
* 不堆砌术语
* 每个概念都尽量配代码
* 解释要深入浅出
* 对容易误解的地方要重点提醒
* 要有工程味，不要像学校教材
* 多使用「Java 中你可能会这样做，但 Go 中更推荐这样做」的表达方式

---

## 十五、执行方式

请你直接在当前目录中创建完整项目。

执行顺序：

1. 创建目录结构
2. 生成根目录文档
3. 生成每个 lesson 的 README.md
4. 为每个 lesson 添加可运行 examples
5. 生成 exercises
6. 生成 projects
7. 生成 cheat-sheets
8. 运行 gofmt
9. 尽可能运行 go test
10. 最后输出执行总结

---

## 十六、验收标准

最终项目必须满足：

* 目录结构完整
* 每个 lesson 都有 README
* 每个 lesson 都有代码示例
* 至少 16 个 lessons
* 至少 55 道练习题
* 至少 4 个实战项目
* 至少 4 个 cheat sheets
* 文档面向 Java 开发者
* 代码可以运行
* 项目具备自学价值
* 不是空壳，不是提纲，而是可直接学习的课程材料

---

## 十七、最后输出

完成后请输出：

```text
UltraWork 完成报告

1. 已创建的目录
2. 已生成的 Lessons
3. 已生成的练习题数量
4. 已生成的项目
5. 已生成的速查表
6. 已执行的格式化/测试命令
7. 可能存在的限制或后续优化建议
8. 推荐学习者如何开始第一天学习
```

现在开始执行。

```
```

