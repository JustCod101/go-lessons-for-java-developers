# Go Lessons for Java Developers

> 如果你会 Java，想快速掌握 Go，这套课程是为你设计的。

本课程不走传统教材路线，而是以 **Java 工程师的视角** 出发，通过对比、类比、踩坑提醒来让你理解 Go 的设计哲学。每一节都有可运行的代码，每一章都有针对性练习，最终通过 4 个实战项目把你从 "看过 Go 语法" 变成 "能写 Go 代码"。

---

## 这套课程适合谁

| 前提 | 说明 |
|------|------|
| ✅ 会 Java 基础语法 | 变量、控制流、面向对象、异常这些你懂 |
| ✅ 了解 Java 后端基础 | 集合、线程、JVM、Spring Boot 有过实践经验更佳 |
| ❌ Go 零基础 | 你不需要会 Go，但需要会 Java |
| ✅ 想快速从 Java 转向 Go | 目标是能写生产级 Go 代码，不只是过语法 |

如果你已经写了一些 Go 但总觉得 "哪里不对劲"，这套课程也能帮你找到思维盲区。

---

## 学完后你能获得什么能力

**目标水平：Go 后端实习 / 初级工程师**

- 能独立阅读 Go 项目源码，理解目录结构和包布局
- 能用 Go 写 CLI 工具、HTTP Server、REST API
- 掌握 goroutine / channel / context，能写并发安全的代码
- 会写单元测试、表驱动测试、benchmark
- 能用 Go 完成数据库 CRUD 操作
- 理解 Go 的工程哲学：简单、直接、组合优于继承

---

## 学习路线（4 周速成）

```
第 1 周 │ Go 环境 → 基础语法 → 函数 → 错误处理 → struct
第 2 周 │ interface → 集合 → 包结构 → 测试 → 标准库
第 3 周 │ goroutine → channel → context → HTTP 服务 → 数据库 CRUD
第 4 周 │ Task API 项目 → 内存缓存项目 → 并发爬虫项目 → 面试复盘
```

详见 [roadmap.md](roadmap.md)

---

## 如何使用本课程

**正确姿势：**

1. **先读 README**：每节 Lesson 的 README 会先讲"这是什么 + 和 Java 的对比"
2. **再跑示例代码**：`go run` 亲手跑一遍，不要只看
3. **改 3 次代码**：这是最小要求，改得越多理解越深
4. **做练习题**：每节的练习题是强制完成项，不是可选项
5. **做项目**：4 周时的综合实战，是你真正掌握的证明

**错误姿势：**

- 只看不练（等于没学）
- 跳过练习题（练习题是帮你发现理解漏洞的）
- 跳过项目（项目是真正的能力证明）

---

## 推荐学习节奏

| 周期 | 内容 | 每天耗时 |
|------|------|----------|
| 第 1-2 周 | 语言基础（Lesson 01-11） | 2-3 小时 |
| 第 3 周 | 并发 + Web（Lesson 08-13） | 2-3 小时 |
| 第 4 周 | 4 个项目 + 复盘 | 3-4 小时 |

如果时间紧张，可以压缩到 6-8 周，但 **不建议拉得更长**——拖延会让你遗忘前面的内容。

---

## 课程目录

```
go-lessons-for-java-developers/
├── lessons/
│   ├── lesson-01-go-overview/           Go 语言整体认知
│   ├── lesson-02-environment-and-modules/  环境与模块
│   ├── lesson-03-basic-syntax/          基础语法
│   ├── lesson-04-functions-and-error-handling/  函数与错误处理
│   ├── lesson-05-struct-method-interface/  struct / method / interface
│   ├── lesson-06-collections/            集合（slice / map）
│   ├── lesson-07-packages-and-project-layout/  包与项目结构
│   ├── lesson-08-goroutine-and-channel/  goroutine 和 channel
│   ├── lesson-09-context-and-concurrency-control/  context 与并发控制
│   ├── lesson-10-standard-library/      标准库
│   ├── lesson-11-testing/               测试
│   ├── lesson-12-http-server/           HTTP 服务
│   ├── lesson-13-database-crud/         数据库 CRUD
│   ├── lesson-14-mini-project-task-api/  综合项目：Task API
│   ├── lesson-15-mini-project-cache/    综合项目：内存缓存
│   └── lesson-16-next-step/             进阶路线
├── exercises/          55+ 道分级练习题
├── projects/           4 个实战项目
└── cheat-sheets/       4 份速查表
```

---

## 项目实战

完成语言基础后，你将独立完成 4 个项目：

| 项目 | 技术栈 | 难度 |
|------|--------|------|
| **cli-todo** | flag, JSON 文件持久化 | ⭐⭐ 入门 |
| **http-task-api** | net/http, JSON, middleware | ⭐⭐⭐ 中等 |
| **memory-cache** | sync.Mutex, TTL, benchmark | ⭐⭐⭐ 中等 |
| **concurrent-crawler** | goroutine, channel, worker pool | ⭐⭐⭐⭐ 进阶 |

每个项目都有明确的目标、可运行的代码、以及设计权衡说明。完成后你的简历上可以写：_"用 Go 实现了 4 个项目，包括 CLI 工具、HTTP API、内存缓存和并发爬虫"_。

---

## Java 开发者学习 Go 的核心转变

如果你只记住一件事：**Go 不是 Java**

| Java 思维 | Go 思维 |
|-----------|---------|
| 类继承，OOP 套路 | 组合优于继承，struct + method |
| 一切皆对象 | 不是一切皆对象，函数也是一等公民 |
| try-catch 异常控制流 | 显式 error 返回，不推荐用异常做控制流 |
| 接口需要显式声明实现 | 接口隐式实现，京不要求声明 |
| Thread / ExecutorService | goroutine + channel，更轻量 |
| Maven / Gradle 构建 | Go Modules，更简单 |
| 强类型 + 泛型 | 轻量泛型（Go 1.18+），但别滥用 |
| Spring Boot 大包大揽 | 标准库优先，不过度框架化 |

---

## 常见问题

**Q: Go 和 Java 比，谁更好？**
没有绝对的好坏。Go 擅长并发、网络服务、基础设施；Java 擅长复杂业务逻辑、生态系统。选择取决于场景。

**Q: 学完这套课能找到 Go 工作吗？**
如果你认真做完了所有练习和项目，对 Go 面试题有准备，找到初级 Go 后端工作是可行的。但这只是起点，持续写代码才是关键。

**Q: 需要先安装什么？**
Go 1.22+。具体安装方法见 [Lesson 02](../lessons/lesson-02-environment-and-modules/README.md)。

**Q: 看不懂某些概念怎么办？**
先跑代码。Go 的设计哲学是"运行比理解更重要"——先让它跑起来，然后慢慢琢磨为什么。

---

## 学习资源

- [Go 官方文档](https://go.dev/doc/)
- [Go 语言圣经（中文）](https://github.com/golang-china/gopl-zh)
- [Go by Example](https://gobyexample.com/)
- [Effective Go（官方）](https://go.dev/doc/effective_go)

---

## 开始学习

从 [Lesson 01: Go 语言整体认知](../lessons/lesson-01-go-overview/README.md) 开始。