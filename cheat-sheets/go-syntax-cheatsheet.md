# Go 语法速查表

> 给 Java 开发者的 Go 语法对照速查。左侧是 Java，右侧是 Go。

---

## 变量与常量

| Java | Go | 说明 |
|------|-----|------|
| `int x = 10;` | `x := 10` | 类型推断 |
| `int x;` | `var x int` | 声明变量 |
| `final int X = 10;` | `const X = 10` | 常量 |
| `String name = "Alice";` | `name := "Alice"` | 字符串 |
| `boolean flag = true;` | `flag := true` | 布尔 |
| `double score = 95.5;` | `score := 95.5` | 浮点数 |

---

## 基本类型

| Java | Go | 说明 |
|------|-----|------|
| `int` | `int` | 平台相关 |
| `long` | `int64` | 64位整数 |
| `double` | `float64` | 64位浮点 |
| `boolean` | `bool` | 布尔 |
| `String` | `string` | 字符串（不可变） |
| `char` | `rune` | Unicode 码点 |
| `byte[]` | `[]byte` | 字节切片 |

---

## 类型转换

| Java | Go |
|------|-----|
| `(int) 10.9` | `int(10.9)` |
| `String.valueOf(123)` | `strconv.Itoa(123)` |
| `Integer.parseInt("123")` | `strconv.Atoi("123")` |

---

## 函数

```go
// 单返回值
func add(a, b int) int {
    return a + b
}

// 多返回值
func div(a, b int) (int, int) {
    return a / b, a % b
}

// 命名返回值
func sum(a, b, c int) (result int) {
    result = a + b + c
    return // 不用写 return result
}

// 匿名函数
add := func(a, b int) int { return a + b }
```

---

## struct

```go
// 定义
type Person struct {
    Name string
    Age  int
}

// 工厂函数（代替构造函数）
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}

// 初始化
p := Person{Name: "Alice", Age: 30}
p2 := &Person{Name: "Bob", Age: 25}
```

---

## method

```go
// value receiver（不会修改原值）
func (p Person) String() string {
    return fmt.Sprintf("Person{Name:%s, Age:%d}", p.Name, p.Age)
}

// pointer receiver（可以修改原值）
func (p *Person) Birthday() {
    p.Age++
}
```

---

## interface

```go
// 定义
type Reader interface {
    Read(p []byte) (n int, err error)
}

// 隐式实现（不需要 implements 关键字）
type MyReader struct{}
func (r MyReader) Read(p []byte) (n int, err error) {
    // 实现...
    return len(p), nil
}

// 空接口（任何类型）
var i any = 42
var s any = "hello"
```

---

## 错误处理

```go
// 返回 error
func readFile(name string) ([]byte, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, err // 向上传递
    }
    return data, nil
}

// 检查 error
data, err := readFile("test.txt")
if err != nil {
    fmt.Println("错误:", err)
    return
}

// 包装错误
return nil, fmt.Errorf("copy failed: %w", err)

// 判断错误类型
if errors.Is(err, os.ErrNotExist) { ... }
if errors.As(err, &myErr) { ... }
```

---

## 控制流

```go
// if（条件不加括号）
if x := 10; x > 5 {
    fmt.Println("x > 5")
}

// for（Go 唯一的循环关键字）
for i := 0; i < 10; i++ { ... }    // 传统 for
for i < 10 { ... }                 // while 替代
for { ... }                        // 无限循环

// switch（不用 break，自动 break）
switch x {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
default:
    fmt.Println("other")
}

// fallthrough 穿透
switch x {
case 1:
    fmt.Println("one")
    fallthrough
case 2:
    fmt.Println("two or one")
}
```

---

## slice

```go
// 创建
s := []int{1, 2, 3}
s := make([]int, 0)       // empty slice
s := make([]int, 5)       // len=5, 零值
s := make([]int, 0, 10)   // len=0, cap=10

// 添加元素
s = append(s, 4)

// 遍历
for i, v := range s {
    fmt.Println(i, v)
}

// 切片（共享底层数组）
t := s[1:3] // t[0]=s[1], t[1]=s[2]

// len 和 cap
fmt.Println(len(s)) // 元素数量
fmt.Println(cap(s)) // 底层数组容量
```

---

## map

```go
// 创建
m := map[string]int{"Alice": 30, "Bob": 25}
m := make(map[string]int)

// 添加/获取
m["Charlie"] = 35
age := m["Alice"]

// 检查存在
age, ok := m["David"] // ok=false 表示不存在

// 删除
delete(m, "Bob")

// 遍历
for k, v := range m {
    fmt.Println(k, v)
}
```

---

## goroutine

```go
// 启动 goroutine
go func() {
    fmt.Println("异步执行")
}()

// 匿名函数带参数
go func(msg string) {
    fmt.Println(msg)
}("hello")
```

---

## channel

```go
// 创建
ch := make(chan int)           // unbuffered
ch := make(chan int, 3)        // buffered

// 发送
ch <- 42

// 接收
v := <-ch

// 关闭
close(ch)

// 遍历
for v := range ch {
    fmt.Println(v)
}

// select 多路复用
select {
case v := <-ch1:
    fmt.Println("ch1:", v)
case ch2 <- 100:
    fmt.Println("sent to ch2")
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

---

## context

```go
// 创建带超时
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// 在请求中传递
req = req.WithContext(ctx)

// 检查取消
select {
case <-ctx.Done():
    fmt.Println("取消:", ctx.Err()) // context.DeadlineExceeded
    return
default:
    // 继续执行
}

// 传递值
ctx = context.WithValue(ctx, "userID", 123)
userID := ctx.Value("userID").(int)
```

---

## sync 包

```go
// WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()

// Mutex（互斥锁）
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()

// RWMutex（读写锁，读多写少时用）
var mu sync.RWMutex
mu.RLock()    // 读
fmt.Println(m["key"])
mu.RUnlock()

// Once（只执行一次）
var once sync.Once
once.Do(func() { fmt.Println("只执行一次") })

// Atomic
var n int64
atomic.AddInt64(&n, 1)
atomic.LoadInt64(&n)
```

---

## defer

```go
// defer 在函数退出时执行，多个 defer 按栈顺序
func read() {
    f, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer f.Close() // 函数退出时执行

    // 读取文件
}
```

---

## 测试

```go
// 测试文件：xxx_test.go
import "testing"

func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}

// 表驱动测试
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"2+3", 2, 3, 5},
        {"0+0", 0, 0, 0},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := add(tt.a, tt.b); got != tt.want {
                t.Errorf("add() = %v, want %v", got, tt.want)
            }
        })
    }
}

// Benchmark
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(2, 3)
    }
}
```

---

## 指针

```go
x := 10
p := &x        // *int，指向 x 的地址
fmt.Println(*p) // 10

*p = 20        // 通过指针修改 x
fmt.Println(x) // 20
```

---

## 结构体标签

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age,omitempty"`
}

// JSON 编解码
data, _ := json.Marshal(p)           // {"name":"Alice","age":30}
json.Unmarshal(data, &p)
```

---

## package import

```go
import (
    "fmt"           // 标准库
    "os"            // 标准库
    "myapp/internal" // 相对路径
    mymodule "github.com/user/mymodule" // 别名
)
```

---

## 可见性规则

- **大写开头**：导出（public）
- **小写开头**：不导出（private）

```go
type Person struct { // 导出
    Name string     // 导出字段
    age int         // 不导出字段
}

func hello() { } // 不导出
func Hello() { } // 导出
```

---

## panic 与 recover

```go
// panic 触发崩溃（尽量少用）
panic("发生了严重错误")

// recover 在 defer 中捕获 panic
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("恢复:", r)
        }
    }()
    panic("oops")
}
```

---

## 空值（nil）

| Java | Go |
|------|-----|
| `null` | `nil`（pointer, channel, slice, map, interface, func） |

```go
var p *int = nil      // nil pointer
var s []int = nil     // nil slice
var m map[string]int = nil // nil map
var c chan int = nil  // nil channel
var i any = nil       // nil interface
```