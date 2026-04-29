# Java 到 Go 映射速查

> 用你已经熟悉的 Java 概念，快速理解对应的 Go 实现。

---

## 类型系统

| Java | Go | 说明 |
|------|-----|------|
| `class` | `struct` | Go 没有 class，用 struct 定义数据结构 |
| `extends` | 组合（嵌入） | Go 用 struct 嵌入实现组合，没有继承 |
| `implements` | 隐式实现 | Go 接口不需要声明实现，只要方法匹配 |
| `interface Foo {}` | `interface{}` 或 `any` | 空接口 |
| `Object` | `any` | 空接口可以存储任何值 |
| `List<T>` | `[]T`（slice） | 动态数组 |
| `ArrayList<T>` | slice + append | 更底层的实现 |
| `LinkedList<T>` | 无直接对应 | Go 没有内置链表 |
| `Map<K, V>` | `map[K]V` | 内置 hash map |
| `HashMap<K, V>` | `map[K]V`（无序） | Go 的 map 不保证顺序 |
| `HashSet<T>` | `map[T]struct{}` | 需要自实现 Set |
| `Queue<T>` | `chan T` | Go 用 channel 实现队列 |
| `Stack<T>` | `[]T` | 用 slice 的 append/pop 模拟 |
| `int` | `int` | 平台相关（32/64 位） |
| `long` | `int64` | 明确 64 位 |
| `boolean` | `bool` | 拼写不同 |
| `String` | `string` | 不可变字符串 |
| `void` | 无返回值 | 用空 return 或省略 |

---

## 面向对象

| Java | Go | 说明 |
|------|-----|------|
| `public class Foo` | `type Foo struct` | Go 没有 class |
| 构造函数 | 工厂函数 `NewXxx()` | Go 没有构造函数 |
| `this` | receiver | 方法的第一个参数 |
| `super` | 无 | Go 没有继承 |
| 静态方法 | 包级函数 | Go 没有 static |
| `final` | 无对应 | 用 const 或不导出 |
| `abstract` | interface | 更轻量的接口 |
| `enum` | const + iota | 枚举的实现方式 |

---

## 函数

| Java | Go | 说明 |
|------|-----|------|
| `void foo()` | `func foo()` | 无返回值 |
| `int foo()` | `func foo() int` | 单返回值 |
| 无多返回值 | `func foo() (int, error)` | Go 的多返回值 |
| 匿名内部类 | 匿名函数 + closure | 函数是一等公民 |
| `Function<T,R>` | `func(T) R` | 函数类型内置 |
| `Consumer<T>` | `func(T)` | 消费型函数 |

---

## 错误处理

| Java | Go | 说明 |
|------|-----|------|
| `try-catch` | `if err != nil` | Go 没有异常机制 |
| `throw` | `panic` | panic 是严重错误，不是正常流程 |
| `throws` | `error` 返回值 | 调用者必须处理 |
| `Exception` | `error` 接口 | 只是一个接口 |
| `RuntimeException` | `panic` | 严重错误 |
| `finally` | `defer` | 函数退出时执行 |

---

## 并发

| Java | Go | 说明 |
|------|-----|------|
| `Thread` | `go func()` | goroutine 更轻量 |
| `Runnable` | `func()` | 匿名函数是入口 |
| `ExecutorService` | `goroutine + channel` | Go 的并发哲学不同 |
| `Future<T>` | `<-chan T` | channel 是异步结果 |
| `CompletableFuture` | `goroutine + channel` | 组合多个异步操作 |
| `synchronized` | `sync.Mutex` | 互斥锁 |
| `ReentrantLock` | `sync.RWMutex` | 读写锁 |
| `CountDownLatch` | `sync.WaitGroup` | 等待一组任务 |
| `BlockingQueue` | `chan` | Go 的 channel |
| `Semaphore` | `chan` 或 `semaphore` 包 | 信号量 |
| `ThreadLocal` | 包级变量 | Go 不鼓励 ThreadLocal |
| `volatile` | `sync/atomic` | 原子操作 |

---

## 集合框架

| Java | Go | 说明 |
|------|-----|------|
| `List<String>` | `[]string` | slice |
| `ArrayList<String>` | slice | 用 append 添加 |
| `HashMap<String, Int>` | `map[string]int` | 内置 map |
| `HashSet<T>` | `map[T]struct{}` | 自实现 Set |
| `Iterator` | `for range` | 遍历的方式 |
| `stream().filter()` | 手写循环 | Go 没有 Stream API |
| `Collections.sort()` | `sort.Slice()` | 排序函数 |
| `Arrays.asList()` | `[]string{...}` | 切片字面量 |

---

## 包与可见性

| Java | Go | 说明 |
|------|-----|------|
| `package com.example` | `package example` | Go 包名简短 |
| `import com.example.Foo` | `import "example"` | import 是完整路径 |
| `public` | 大写开头 | 导出标识符 |
| `private` | 小写开头 | 不导出标识符 |
| `protected` | 无 | Go 没有 protected |
| 包内可见 | 默认 | 小写即包内可见 |

---

## 控制流

| Java | Go | 说明 |
|------|-----|------|
| `if (x > 5)` | `if x > 5` | 条件不加括号 |
| `for (int i=0; i<n; i++)` | `for i := 0; i < n; i++` | 略有不同 |
| `while (condition)` | `for condition {}` | Go 没有 while |
| `for (item : list)` | `for i, v := range list {}` | range 遍历 |
| `switch (value)` | `switch value {}` | 更强大的 switch |
| `break` | `break` | 相同 |
| `continue` | `continue` | 相同 |

---

## 测试

| Java | Go | 说明 |
|------|-----|------|
| JUnit | `testing` 包 | Go 内置测试框架 |
| `@Test` | `func TestXxx(t *testing.T)` | 命名约定 |
| `assertEquals` | `t.Errorf` | 断言方式不同 |
| `@BeforeEach` | 初始化在测试函数内 | 没有 suite |
| `@ParameterizedTest` | 手写循环 | 没有参数化测试 |
| `@Disabled` | `t.Skip()` | 跳过测试 |
| `Mockito` | 手动 mock / 接口 | 没有官方 mock |
| `@Nested` | 子测试 `t.Run` | 嵌套测试 |

---

## 构建工具

| Java | Go | 说明 |
|------|-----|------|
| Maven/Gradle | Go Modules | 依赖管理 |
| `pom.xml` | `go.mod` | 依赖文件 |
| `mvn package` | `go build` | 编译 |
| `mvn test` | `go test` | 测试 |
| `mvn install` | `go install` | 安装二进制 |
| 私有仓库 | `go mod private` | 私有模块 |
| Maven Central | proxy.golang.org | Go 的模块代理 |

---

## Web 框架

| Java | Go | 说明 |
|------|-----|------|
| Spring Boot | `net/http`（标准库） | Go 鼓励先用标准库 |
| `@RestController` | `http.HandlerFunc` | 更底层的处理方式 |
| `@RequestMapping` | `http.HandleFunc` | 路由定义 |
| `@PathVariable` | 手动解析 URL | 参数解析方式 |
| `@RequestBody` | `json.NewDecoder(r.Body)` | 请求体解析 |
| `ResponseEntity` | `w.Write` + JSON | 响应方式 |
| Filter/Interceptor | middleware | 中间件模式类似 |
| Tomcat/Jetty | `net/http` 内置 | Go 内置 HTTP 服务器 |

---

## 数据库

| Java | Go | 说明 |
|------|-----|------|
| JDBC | `database/sql` | 类似的设计 |
| `DataSource` | `sql.DB` | 连接池代表 |
| `ResultSet` | `sql.Rows` | 查询结果 |
| MyBatis | `sqlx` / `GORM` | ORM 库 |
| JPA/Hibernate | `GORM` / `ent` | ORM 框架 |
| 事务 | `db.Begin()` | 事务处理 |
| SQL 注入防御 | 参数化查询 | Go 原生支持 |

---

## 项目结构

| Java | Go | 说明 |
|------|-----|------|
| 分层架构 | 扁平结构 | Go 不鼓励过度分层 |
| `controller/service/repository` | `handler/service/repository` | 可以用类似的结构 |
| `entity` | `model` / `domain` | 数据模型 |
| `config` | `internal/config` | 配置 |
| `exception` | `error` | Go 不推荐异常 |

---

## 常用的对应关系速记

```
Java → Go
class → struct
extends → 组合（嵌入）
implements → 隐式接口实现
Exception → error
try-catch → if err != nil
throw → panic（慎用）
Thread → goroutine
ExecutorService → goroutine + channel + sync
synchronized → sync.Mutex
List → slice
Map → map
HashSet → map[T]struct{}
null → nil
toString() → fmt.Sprintf("%+v", x)
@Override → 无标签，靠命名
Maven → Go Modules
JUnit → testing
Stream API → 手写循环
```

---

## 记住的关键转变

1. **Go 不是面向对象的**——不要用 Java OOP 思维套 Go
2. **error 是返回值**——不是异常，是值
3. **goroutine 很轻量**——大胆用，不需要线程池
4. **接口是隐式实现**——不需要声明
5. **组合优于继承**——用嵌入代替继承