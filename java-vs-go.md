# Java 与 Go：核心差异对照表

> 本文档帮助 Java 开发者快速建立 Go 的心智模型。对比不是为了让两者对立，而是帮助你把已有的 Java 知识迁移到 Go，同时提醒你哪些 Java 习惯需要调整。

---

## 语言哲学对比

| 维度 | Java | Go |
|------|------|-----|
| **设计目标** | 面向对象企业级应用 | 简单、快速、并发 |
| **类型系统** | 强类型 + 继承层次 | 强类型 + 组合优于继承 |
| **内存管理** | JVM GC | 逃逸分析 + 轻量 GC |
| **并发模型** | Thread / ExecutorService | goroutine + channel |
| **错误处理** | Exception 控制流 | 显式 error 返回值 |
| **编译速度** | 较慢（Maven/Gradle） | 极快 |
| **构建工具** | Maven / Gradle | Go Modules |
| **运行方式** | JVM 字节码 | 本地机器码 |
| **适用场景** | 复杂业务逻辑、企业系统 | 网络服务、基础设施、并发处理 |

---

## 类型系统

| Java | Go | 说明 |
|------|-----|------|
| `class` | `struct` | Go 的 struct 没有构造函数，没有继承 |
| `class Foo implements Bar` | `func (f Foo) Bar() {}` | Go 接口隐式实现，无需声明 |
| `interface Foo {}` | `interface{}` 或 `any` | 空接口相当于 Java 的 `Object` |
| `List<T>` | `slice` | Go 的 slice 是动态数组，不是链表 |
| `Map<K, V>` | `map` | 内置 map，相当于 HashMap |
| `HashSet<T>` | 自实现 `map[T]struct{}` | Go 没有内置 Set |
| `int` / `long` | `int` / `int64` | Go 的 int 平台相关（32/64 位） |
| `boolean` | `bool` | 拼写不同 |
| `String` | `string` | 不可变字符串 |
| `void` | 无返回值函数 | 用空 `return` 或直接省略 |

---

## 面向对象

| Java 概念 | Go 对应 | 关键差异 |
|-----------|--------|----------|
| class | struct + method | Go 没有类，没有继承，没有方法重写 |
| extends | 组合（struct 内嵌） | Go 用组合代替继承 |
| implements | 隐式实现 | 不需要显式声明 |
| 抽象类 | 空 interface | Go 的接口更轻量 |
| 静态方法 | 包级函数 | Go 没有 static 关键字 |
| 构造函数 | 工厂函数或 NewXxx 函数 | Go 没有构造函数，用 `NewXxx()` 代替 |
| this | receiver | method 的第一个参数是 receiver |
| super | 无 | Go 没有 super，因为没有继承 |

**Java 中的继承：**
```java
class Animal {
    void speak() { }
}

class Dog extends Animal {
    @Override
    void speak() {
        System.out.println("Woof");
    }
}
```

**Go 中的组合：**
```go
type Animal struct {
    Name string
}

func (a *Animal) Speak() {
    fmt.Println("...")
}

type Dog struct {
    Animal // 组合，相当于继承
}

func (d *Dog) Speak() {
    fmt.Println("Woof")
}
```

---

## 函数

| Java | Go | 说明 |
|------|-----|------|
| `void foo()` | `func foo()` | 无返回值函数 |
| `int foo()` | `func foo() int` | 单返回值 |
| 无多返回值 | `func foo() (int, error)` | Go 支持多返回值 |
| 无命名返回值 | `func foo() (result int)` | Go 支持命名返回值 |
| lambda / 匿名类 | 匿名函数 + closure | Go 的函数是一等公民 |
| `Function<T, R>` | `func(T) R` | Go 的函数类型是内置的 |

---

## 错误处理

| Java | Go | 说明 |
|------|-----|------|
| `try-catch` | 显式 error 返回 | Go 没有异常机制 |
| `throw new Exception()` | `return errors.New("...")` | 错误是值，不是异常 |
| `throws` 声明 | `error` 作为返回值 | 调用者必须处理 error |
| `Exception` 层次 | `error` 接口 | Go 的 error 只是一个接口 |
| `RuntimeException` | `panic` | 严重错误用 panic，但不要滥用 |
| `finally` | `defer` | defer 在函数退出时执行 |

**Java 的异常处理：**
```java
try {
    String result = readFile();
    process(result);
} catch (IOException e) {
    handle(e);
} finally {
    close();
}
```

**Go 的错误处理：**
```go
result, err := readFile()
if err != nil {
    handle(err)
    return
}
err = process(result)
if err != nil {
    handle(err)
    return
}
defer close()
```

**Go 的 error wrapping：**
```go
_, err := io.Copy(dst, src)
if err != nil {
    return fmt.Errorf("copy failed: %w", err)
}
```

---

## 并发

| Java | Go | 说明 |
|------|-----|------|
| `Thread` | `go func()` | goroutine 是轻量级线程 |
| `Runnable` | `func()` | 匿名函数作为 goroutine 入口 |
| `ExecutorService` | `goroutine + channel + sync` | Go 的并发哲学不同 |
| `synchronized` | `sync.Mutex` | 互斥锁 |
| `ReentrantLock` | `sync.RWMutex` | 读写锁 |
| `CountDownLatch` | `sync.WaitGroup` | 等待一组任务完成 |
| `BlockingQueue` | `chan` | channel 是 Go 的并发核心 |
| `Semaphore` | `chan` 或 `golang.org/x/sync/semaphore` | 信号量 |
| `ThreadLocal` | 包级变量 + 手动传递 | Go 不鼓励使用 ThreadLocal |
| `CompletableFuture` | `goroutine + channel` | Go 用管道组合异步任务 |
| `ExecutorService.submit()` | `make(chan result)` | 获取结果的方式不同 |

**Java Thread：**
```java
ExecutorService executor = Executors.newFixedThreadPool(4);
Future<String> future = executor.submit(() -> {
    return fetchData();
});
String result = future.get();
```

**Go goroutine：**
```go
resultCh := make(chan string)
go func() {
    resultCh <- fetchData()
}()
result := <-resultCh
```

---

## 集合框架

| Java | Go | 说明 |
|------|-----|------|
| `List<String>` | `[]string` (slice) | 动态数组 |
| `ArrayList<String>` | slice + append | Go 的 slice 更底层 |
| `LinkedList<String>` | 无直接对应 | Go 没有内置链表 |
| `Map<String, Integer>` | `map[string]int` | 内置 map |
| `HashMap` | `map`（无序） | Go 的 map 无序 |
| `HashSet<T>` | `map[T]struct{}` | 需要自实现 Set |
| `Queue` | slice + channel | Go 用 channel 做队列 |
| `Stack` | slice | 用 append/pop 模拟 |
| `Arrays.asList()` | `[]string{...}` | 切片字面量 |
| `stream()` | 无直接对应 | Go 没有 Stream API |

---

## 包与可见性

| Java | Go | 说明 |
|------|-----|------|
| `package com.example` | `package example` | Go 包名简短 |
| `import com.example.Foo` | `import "example"` | Go import 是完整路径 |
| `public` | 大写开头 | 导出标识符用大写 |
| `private` | 小写开头 | 未导出标识符用小写 |
| `protected` | 无 | Go 没有 protected |
| `default`（包内） | 小写标识符 | Go 的默认可见性就是包内 |
| `import java.util.*` | `import ("fmt"; "strings")` | Go 的 import 更明确 |

---

## 控制流

| Java | Go | 说明 |
|------|-----|------|
| `if (condition) {}` | `if condition {}` | Go 条件不加括号 |
| `for (int i=0; i<n; i++)` | `for i := 0; i < n; i++` | 语法略有不同 |
| `while (condition)` | `for condition {}` | Go 没有 while 关键字 |
| `for (item : collection)` | `for i, v := range collection {}` | range 遍历 |
| `switch (value)` | `switch value {}` | Go 的 switch 更强大 |
| `break` | `break` | 相同 |
| `continue` | `continue` | 相同 |

---

## 测试

| Java | Go | 说明 |
|------|-----|------|
| JUnit | `testing` 包 | Go 内置测试框架 |
| `@Test` | `func TestXxx(t *testing.T)` | 测试函数命名规则 |
| `assertEquals` | 直接比较或 `t.Errorf` | Go 测试风格不同 |
| `@BeforeEach` | `TestXxx` 内初始化 | Go 不支持 suite |
| `Mockito` | 手动 mock 或接口 | Go 没有官方 mock 框架 |
| `assertThat` | 表驱动测试 | Go 偏好表驱动测试 |
| `@Ignore` | `t.Skip()` | 跳过测试的方式 |
| JUnit 5 parameterized | 手写循环 | Go 没有内置参数化测试 |

---

## 构建工具对比

| Java | Go | 说明 |
|------|-----|------|
| `mvn package` | `go build` | 编译输出 |
| `mvn test` | `go test` | 运行测试 |
| `mvn clean` | `go clean` | 清理构建产物 |
| `pom.xml` | `go.mod` | 依赖管理文件 |
| `mvn dependency:tree` | `go mod graph` | 查看依赖关系 |
| Gradle Kotlin DSL | 无，go.mod 很简单 | Go Modules 更轻量 |
| Maven central | proxy.golang.org | Go 的模块代理 |
| `mvn install` | `go install` | 安装二进制 |

---

## Web 框架对比

| Java | Go | 说明 |
|------|-----|------|
| Spring Boot | `net/http`（标准库） | Go 鼓励先用标准库 |
| `@RestController` | `http.HandlerFunc` | Go 的 HTTP 处理更底层 |
| `@RequestMapping` | `http.HandleFunc` | 路由定义方式不同 |
| `@PathVariable` | `url path parsing` | 参数解析方式 |
| `@RequestBody` | `json.NewDecoder` | 请求体解析 |
| `@ResponseBody` | `json.NewEncoder` | 响应体编码 |
| Filter / Interceptor | middleware | 中间件模式相似 |
| Tomcat / Undertow | `net/http` 默认服务器 | Go 内置 HTTP 服务器 |
| Spring Data JPA | `database/sql` / GORM | 数据库访问方式 |

---

## 项目结构对比

| Java 多层架构 | Go 风格 | 说明 |
|---------------|--------|------|
| `com.example.controller` | `handler` / `http` | Controller 层 |
| `com.example.service` | `service` | Service 层（Go 项目有时简化掉） |
| `com.example.repository` | `repository` / `store` | 数据访问层 |
| `com.example.entity` | `model` / `domain` | 数据模型 |
| `com.example.config` | `internal` / `config` | 配置 |
| `com.example.exception` | 直接返回 error | Go 不推荐异常 |

**Go 推荐的简单结构：**
```
myproject/
├── cmd/
│   └── myapp/
│       └── main.go          # 入口
├── internal/
│   ├── handler/              # HTTP handlers
│   ├── service/              # 业务逻辑
│   └── repository/           # 数据访问
├── pkg/                      # 可导出给其他项目用的包
└── go.mod
```

**Go 反对过度分层**——如果你的项目只有 handler + service + repository 三层，Go 的风格是让你把 service 这一层根据需要合并或去掉。

---

## 常见的 Java 习惯需要改

| Java 习惯 | Go 正确做法 | 为什么 |
|-----------|-------------|--------|
| 到处定义接口 | 按需定义接口 | Go 的接口是行为抽象，不要过度设计 |
| 类继承思维 | 组合思维 | Go 鼓励组合，不鼓励继承 |
| 异常控制流 | 显式 error | Go 的 error 是值类型，不应该用来做控制流 |
| 泛型依赖 | 轻量泛型 | Go 1.18+ 支持泛型，但不要滥用 |
| 框架依赖 | 标准库优先 | Go 的标准库足够强，不过度框架化 |
| 单例模式 | 包级变量 + sync | Go 的包级变量天然单例 |
| 链式调用 | 显式处理 | Go 没有 Builder 链式语法糖 |
| 反射 | 避免反射 | Go 的反射性能差，且不安全 |

---

## 快速对照卡（速记）

| Java | Go | 备注 |
|------|-----|------|
| class | struct | |
| extends | 组合 | |
| implements | 隐式实现 | |
| Exception | error | |
| try-catch | if err != nil | |
| throw | panic（慎用） | |
| Thread | goroutine | |
| ExecutorService | goroutine + channel | |
| synchronized | sync.Mutex | |
| List | slice | |
| Map | map | |
| Set | map[T]struct{} | |
| null | nil | |
| toString() | fmt.Printf("%+v") | |
| @Override | 无标签，靠命名约定 | |
| Maven | Go Modules (go.mod) | |
| JUnit | testing | |
| Stream API | 无，手写循环 | |

---

## 记住最重要的转变

**从 Java 到 Go，关键是记住：**

1. **组合优于继承** —— 不要用继承来复用代码
2. **接口隐式实现** —— 不需要声明，只要方法匹配就行
3. **error 是返回值** —— 不要用 try-catch 的思维处理 error
4. **goroutine 很轻量** —— 放心大胆地用，不需要池
5. **简单优于复杂** —— Go 的哲学是简单、直接、不炫技

当你发现自己写的 Go 代码看起来像 Java 时，停下来问自己：_"这是 Go 风格的吗？"_

---

## 下一步

想了解更多细节？继续学习：
- [Lesson 01: Go 语言整体认知](../lessons/lesson-01-go-overview/README.md)
- [Lesson 05: struct / method / interface](../lessons/lesson-05-struct-method-interface/README.md)
- [Lesson 08: goroutine 和 channel](../lessons/lesson-08-goroutine-and-channel/README.md)