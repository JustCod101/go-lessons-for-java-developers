# Go 学习路线（4 周速成）

> 本路线专为零基础 Go 的 Java 工程师设计。每天 2-4 小时，按顺序推进。

---

## 第 1 周：语言基础（上）

目标是理解 Go 的基本语法和核心设计思路。

| 日期 | 学习内容 | 编码任务 | 复盘问题 |
|------|----------|----------|----------|
| **Day 1** | Lesson 01: Go 语言整体认知 | 跑通 Hello World，理解 Go 的定位 | Go 和 Java 的核心差异是什么？ |
| **Day 2** | Lesson 02: 环境与 Go Modules | 安装 Go，运行 `go mod init`，理解 module vs package | Maven/Gradle 和 Go Modules 的本质区别？ |
| **Day 3** | Lesson 03: 基础语法 | 实现变量、常量、控制流练习 | Go 的零值机制解决了什么问题？ |
| **Day 4** | Lesson 04: 函数与错误处理 | 编写多返回值函数，实现 error 处理 | 为什么 Go 不推荐用 panic 做正常控制流？ |
| **Day 5** | Lesson 05: struct / method / interface | 定义 struct，实现 method，对比 Java class | Go 的隐式接口和 Java 的显式 implements 有什么区别？ |
| **Day 6** | 复习 + 基础练习 | 完成 exercises/basic 前 10 题 | 本周学到的最重要的概念是哪个？ |
| **Day 7** | 休息 | - | - |

**预估耗时：** 每天 2-3 小时

---

## 第 2 周：语言基础（下）

目标是掌握 Go 的集合类型、包结构、测试和标准库。

| 日期 | 学习内容 | 编码任务 | 复盘问题 |
|------|----------|----------|----------|
| **Day 8** | Lesson 06: 集合（slice / map） | 实现一个简化版 ArrayList，对比 Java ArrayList | slice 底层是什么？append 时发生了什么？ |
| **Day 9** | Lesson 07: 包与项目结构 | 搭建一个小型多包项目，理解大小写可见性 | Java 的 package private 和 Go 的小写标识符有什么关系？ |
| **Day 10** | Lesson 11: 测试 | 写表驱动测试，实现 benchmark | Go 测试和 JUnit 的核心差异是什么？ |
| **Day 11** | Lesson 10: 标准库 | 用 fmt/strings/time/os 完成小工具 | 你觉得哪个标准包最实用？为什么？ |
| **Day 12** | 复习 + 中级练习 | 完成 exercises/intermediate 前 10 题 | 哪些 Java 习惯在学习 Go 时需要改变？ |
| **Day 13** | 阶段复盘 | 回头做 Lesson 01-05 的综合练习题 | 第二周和第一周相比，心态有什么变化？ |
| **Day 14** | 休息 | - | - |

**预估耗时：** 每天 2-3 小时

---

## 第 3 周：并发与 Web

目标是掌握 Go 的并发模型和网络编程。

| 日期 | 学习内容 | 编码任务 | 复盘问题 |
|------|----------|----------|----------|
| **Day 15** | Lesson 08: goroutine 和 channel | 实现 worker pool，理解 channel vs BlockingQueue | goroutine 和 Java Thread 的核心区别是什么？ |
| **Day 16** | Lesson 09: context 与并发控制 | 用 context 实现超时控制，对比 Java Future | context 的四种场景各是什么时候用？ |
| **Day 17** | Lesson 12: HTTP 服务 | 用 net/http 实现 REST API，对比 Spring Boot Controller | Go 的 Handler 和 Spring MVC 的 Handler 有什么相似和不同？ |
| **Day 18** | Lesson 13: 数据库 CRUD | 实现 Task 的增删改查，对比 Java JDBC/MyBatis | 连接池、事务、SQL 注入——Go 和 Java 处理方式有何不同？ |
| **Day 19** | HTTP + DB 综合练习 | 把 Task API 连上 SQLite | 为什么 Lesson 13 的设计叫"Repository Pattern"？ |
| **Day 20** | 复习 + 进阶练习 | 完成 exercises/advanced 前 5 题 | 第三周最大的挑战是什么？ |
| **Day 21** | 休息 | - | - |

**预估耗时：** 每天 2-3 小时

---

## 第 4 周：综合项目

目标是独立完成 4 个实战项目，把所有知识点串联起来。

| 日期 | 项目 | 目标 | 验收标准 |
|------|------|------|----------|
| **Day 22** | cli-todo | 实现 todo 命令行工具 | 支持 add / list / done / delete，JSON 持久化 |
| **Day 23** | cli-todo 收尾 + 复习 | 完善测试，修复 bug | go test 通过，go run 正常 |
| **Day 24** | http-task-api | 实现 REST API + middleware | CRUD + 错误处理 + graceful shutdown |
| **Day 25** | http-task-api 收尾 + memory-cache 起步 | 添加单元测试，开始缓存项目 | http-task-api 用 go test 覆盖核心逻辑 |
| **Day 26** | memory-cache | 实现并发安全的内存缓存 + TTL | go test -race 通过，有 benchmark |
| **Day 27** | memory-cache 收尾 + concurrent-crawler 起步 | 添加 eviction 策略，开始爬虫 | 缓存行为符合预期 |
| **Day 28** | concurrent-crawler | 实现并发爬虫 + worker pool | 支持 context cancel、并发控制、去重 |
| **Day 29** | 并发爬虫收尾 | 完成去重 + 输出结果 | 爬虫能正常抓取并输出 |
| **Day 30** | 总复盘 + 面试准备 | 复习 4 周内容，准备面试题 | 能回答常见 Go 面试题 |
| **Day 31** | 总结输出 | 整理项目代码，写学习总结 | 完成 4 个项目 + 学习日志 |

**预估耗时：** 每天 3-4 小时（项目阶段比学习阶段更耗时间）

---

## 每日复盘模板

每晚睡前花 10 分钟回答：

```
1. 今天学了什么？（1-2 句话）
2. 哪个概念最难理解？怎么突破的？
3. 代码写了多少行？改了哪里？
4. 明天要重点解决什么问题？
```

---

## 学习检查点

| 时间 | 检查点 | 通过标准 |
|------|--------|----------|
| 第 1 周结束 | 能跑通 Go 基本语法 | go run 能运行，go test 能通过 |
| 第 2 周结束 | 能读懂 Go 项目结构 | 能用自己的话解释 package/import/module |
| 第 3 周结束 | 能写简单 HTTP 服务 | net/http 能响应 JSON，有单元测试 |
| 第 4 周结束 | 4 个项目全部完成 | 每个项目能 go run 跑通，go test 通过 |

---

## 如果你落后了怎么办

如果某周内容没有按时完成：

1. **不要跳到下一周**——Go 的知识点是递进的，前面没懂后面会更难
2. **减少每天任务量**——如果 Day 15 没完成，Day 16 只做 Day 15 的复习
3. **不要在练习题上糊弄**——如果一道题不懂，花 2 小时研究比跳过更有价值
4. **遇到不懂的查官方文档**——Go 的官方文档是全中文最好懂的编程语言文档

---

## 进阶学习路径（第 5 周起）

完成 4 周基础后，你可以继续深入：

```
继续深入 │
├─ 深入并发：GMP 调度、GC 调优、pprof
├─ Web 框架：Gin / Echo / Fiber
├─ 数据库：GORM / sqlc / ent
├─ 微服务：gRPC、ProtoBuf、Kubernetes client-go
├─ 基础设施：Docker、Kubernetes、云原生
└─ 推荐项目：写一个自己的工具或服务
```

详见 [Lesson 16: 进阶路线](../lessons/lesson-16-next-step/README.md)