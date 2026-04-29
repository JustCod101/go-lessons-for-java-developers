# Go 常见坑速查

> Java 开发者学习 Go 时最容易踩的坑，以及如何避免。

---

## 1. 用 Java OOP 思维硬套 Go

**错误：**
```go
// 试图在 Go 中实现 "继承"
type Dog struct {
    Animal  // 组合，但不是 Java 的继承
    Name    string
}
```

**问题：** Go 没有继承，没有 `extends`，没有 `@Override`。

**正确做法：** 用组合代替继承，理解 Go 的 OOP 是**组合优于继承**。
```go
type Animal struct {
    Age int
}

type Dog struct {
    Animal // 组合，不是继承
    Name   string
}
```

---

## 2. 过度抽象，到处定义 interface

**错误：**
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

func process(r Reader) { ... }
```

**问题：** Java 开发者习惯先定义接口，但 Go 不需要——标准库 concrete type 也可以。

**正确做法：** 按需定义接口，只在需要 mock 或解耦时才定义 interface。

---

## 3. 忽视 error 处理

**错误：**
```go
result, err := doSomething()
fmt.Println(result) // 如果 err 不为 nil，这里可能出问题
```

**问题：** Java 的 Exception 会打断执行，Go 的 error 不处理也不会 panic——但会导致后续问题。

**正确做法：** 养成检查 error 的习惯：
```go
result, err := doSomething()
if err != nil {
    return fmt.Errorf("doSomething failed: %w", err)
}
```

---

## 4. 滥用 panic

**错误：**
```go
if data == nil {
    panic("data is nil")
}
```

**问题：** panic 相当于 Java 的 `RuntimeException`，会导致整个程序崩溃。

**正确做法：** 只在真正无法恢复的错误时使用 panic，普通错误用 error：
```go
if data == nil {
    return errors.New("data is nil")
}
```

---

## 5. goroutine 泄漏

**错误：**
```go
func produce(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    // channel 永远不会关闭，接收方会永远等待
}
```

**问题：** 如果没有接收方，goroutine 会泄漏。

**正确做法：** 确保 channel 正确关闭，或使用 context 取消：
```go
func produce(ctx context.Context, ch chan<- int) {
    for i := 0; i < 10; i++ {
        select {
        case ch <- i:
        case <-ctx.Done():
            return
        }
    }
    close(ch) // 完成后关闭
}
```

---

## 6. channel 乱关闭或忘记关闭

**错误：**
```go
// 关闭已关闭的 channel 会 panic
close(ch)
close(ch) // panic: close of closed channel

// 发送到一个已关闭的 channel 会 panic
close(ch)
ch <- 1 // panic: send on closed channel
```

**正确做法：** 只有发送方关闭 channel，不要在接收方关闭。

---

## 7. slice 共享底层数组

**错误：**
```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]  // b 和 a 共享底层数组
b[0] = 10
fmt.Println(a) // [1 10 3 4 5] - a 被意外修改
```

**问题：** slice 的子切片共享同一个底层数组。

**正确做法：** 如果不想共享，用 `append` 创建新 slice：
```go
b := make([]int, 2)
copy(b, a[1:3])
```

---

## 8. slice append 扩容导致原变量未更新

**错误：**
```go
func appendSlice(s []int) {
    s = append(s, 4) // 扩容后 s 指向新数组，原变量不变
}

func main() {
    s := []int{1, 2, 3}
    appendSlice(s)
    fmt.Println(s) // [1 2 3]，不是 [1 2 3 4]
}
```

**正确做法：** 返回新 slice，或用指针：
```go
func appendSlice(s *[]int) {
    *s = append(*s, 4)
}
```

---

## 9. nil slice vs empty slice

| 类型 | 值 | `len()` | `cap()` | append |
|------|-----|---------|---------|--------|
| `nil slice` | `nil` | 0 | 0 | 正常工作，扩容 |
| `empty slice` | `[]int{}` | 0 | 0 | 正常工作 |

```go
var s []int          // nil slice
s2 := []int{}       // empty slice
s3 := make([]int, 0) // empty slice

// JSON 序列化不同
data, _ := json.Marshal(s)    // null
data2, _ := json.Marshal(s2)  // []
```

---

## 10. map 并发读写 panic

**错误：**
```go
var m = make(map[string]int)

go func() {
    for { m["a"] = 1 }
}()

go func() {
    for { _ = m["a"] }
}()

// 运行时 panic: fatal error: concurrent map read and map write
```

**正确做法：** 用 `sync.Mutex` 或 `sync.RWMutex` 保护，或使用 `sync.Map`：
```go
var mu sync.Mutex
var m = make(map[string]int)

mu.Lock()
m["a"] = 1
mu.Unlock()

mu.RLock()
_ = m["a"]
mu.RUnlock()
```

---

## 11. 不理解 pointer receiver

**错误：**
```go
type Counter struct {
    count int
}

func (c Counter) Increment() {
    c.count++ // 修改的是副本，原对象不变
}

func main() {
    c := Counter{}
    c.Increment()
    fmt.Println(c.count) // 0
}
```

**正确做法：** 如果方法需要修改对象，用 pointer receiver：
```go
func (c *Counter) Increment() {
    c.count++ // 修改原对象
}
```

**经验法则：** 如果不确定，用 pointer receiver——Go 的惯例是偏向指针。

---

## 12. 不会写表驱动测试

**错误：**
```go
func TestAdd(t *testing.T) {
    if add(2, 3) != 5 { t.Error("fail") }
    if add(0, 0) != 0 { t.Error("fail") }
    // ... 更多测试
}
```

**正确做法：** Go 推崇表驱动测试：
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"2+3=5", 2, 3, 5},
        {"0+0=0", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := add(tt.a, tt.b); got != tt.want {
                t.Errorf("add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

---

## 13. 不会使用 context

**错误：**
```go
func longOperation() {
    // 没有超时控制，调用方无法取消
    time.Sleep(10 * time.Second)
}
```

**正确做法：** 使用 context 传递超时/取消信号：
```go
func longOperation(ctx context.Context) error {
    select {
    case <-time.After(10 * time.Second):
        return nil
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

---

## 14. 目录结构过度 Spring 化

**错误：**
```
myproject/
├── controller/
├── service/
├── repository/
├── entity/
├── config/
├── exception/
└── ...
```

**问题：** Go 不需要这么多层级，过于分层会让代码难读。

**正确做法：** Go 项目更扁平：
```
myproject/
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   └── repository/
└── pkg/
    └── util/
```

根据项目规模，甚至可以更简单：
```
myproject/
├── main.go      // 入口
├── handler.go    // HTTP 处理
├── model.go      // 数据模型
└── storage.go    // 存储
```

---

## 15. Go 的 = 不是 ==

**错误：**
```go
if x = 10 { // 这是赋值，不是比较！
    fmt.Println("x 是 10")
}
```

**正确做法：** 用 `==` 比较：
```go
if x == 10 {
    fmt.Println("x 是 10")
}
```

---

## 16. defer 执行时机

**错误：**
```go
func foo() {
    defer fmt.Println("deferred")
    return "early" // 这行不会执行
}
```

**问题：** defer 在 return 之前执行，但 return 语句已经确定了返回值。

**正确做法：** 如果 defer 需要使用返回值，先处理：
```go
func foo() string {
    result := "early"
    defer fmt.Println("deferred, result =", result)
    return result
}
// 输出: deferred, result = early
```

---

## 17. 循环中捕获 goroutine 变量

**错误：**
```go
for _, v := range []int{1, 2, 3} {
    go func() {
        fmt.Println(v) // 所有 goroutine 都打印 3
    }()
}
```

**问题：** 循环变量在所有 goroutine 间共享。

**正确做法：** 传递参数：
```go
for _, v := range []int{1, 2, 3} {
    go func(val int) {
        fmt.Println(val) // 1, 2, 3
    }(v)
}
```

---

## 18. 不理解 nil interface

**错误：**
```go
var err error = nil       // err 是 nil
var r io.Reader = nil     // r 是 nil

var err2 *MyError = nil   // err2 是 nil
```

**问题：** interface 的 nil 判断要小心——如果 interface 持有 nil pointer，它不是 nil。

**正确做法：** 理解 interface 的内部表示（type + value）：
```go
var i any = (*int)(nil)
fmt.Println(i == nil) // false，因为 interface 不是 nil，它有 type 和 value
```

---

## 19. 忘记 close 资源

**错误：**
```go
f, _ := os.Open("file.txt")
// 使用 f
// 忘记 f.Close()
```

**正确做法：** 使用 defer close：
```go
f, err := os.Open("file.txt")
if err != nil {
    return err
}
defer f.Close()
// 使用 f
```

---

## 20. 用 `range` 修改 slice/map

**错误：**
```go
for i, v := range s {
    s[i] = v * 2 // 正确，但有时会忘记
}

for k, v := range m {
    m[k] = v + 1 // 错误：不允许在 range 中修改 map
}
```

**正确做法：** 用索引修改 slice：
```go
for i := range s {
    s[i] *= 2
}
```

---

## 总结：避免踩坑的核心原则

| 原则 | 说明 |
|------|------|
| **不要用 Java 思维套 Go** | Go 的设计哲学和 Java 不同 |
| **error 要检查** | 养成 `if err != nil` 的习惯 |
| **goroutine 要管理** | 用 context、channel、sync 原语 |
| **channel 要正确关闭** | 只让发送方关闭 |
| **map 要加锁** | 普通 map 并发读写会 panic |
| **用 pointer receiver** | 不知道用哪个时，用指针 |
| **写表驱动测试** | Go 推荐的方式 |
| **项目结构要简单** | 不要过度分层 |