# 中级练习

本目录包含 Go 语言中级练习，适合学完前 7 节课程的学习者。

---

## 练习清单

### 1. 实现 Set
**目标**: 理解如何用 map 实现 Set
**输入输出示例**:
```
Set: [1 2 3 5]
包含 3: true
包含 4: false
```
**提示**: Go 没有内置 Set，用 `map[T]struct{}` 实现
**参考解法路径**: 定义 `type Set map[int]struct{}`，添加 `Add`、`Contains` 方法

---

### 2. 实现 Stack
**目标**: 掌握栈的 LIFO 特性
**输入输出示例**:
```
push: 1, 2, 3
pop: 3
pop: 2
栈顶: 1
```
**提示**: 用 slice 实现 push/pop/peek 操作
**参考解法路径**: `type Stack struct { items []int }`，用 append 添加，用 slice 移除

---

### 3. 实现 Queue
**目标**: 掌握队列的 FIFO 特性
**输入输出示例**:
```
入队: 1, 2, 3
出队: 1
出队: 2
队首: 3
```
**提示**: 用 slice 实现，入队用 append，出队用 slice[1:]
**参考解法路径**: `type Queue struct { items []int }`

---

### 4. JSON 编解码
**目标**: 掌握 encoding/json 的使用
**输入输出示例**:
```
原始结构: {Name:Alice Age:30}
JSON 字符串: {"Name":"Alice","Age":30}
反序列化: {Name:Alice Age:30}
```
**提示**: 结构体标签控制 JSON 字段名，`omitempty` 忽略零值
**参考解法路径**: `json.Marshal`、`json.Unmarshal`，结构体标签 `json:"name"`

---

### 5. 读取文件
**目标**: 掌握 os 和 io 包的文件操作
**输入输出示例**:
```
文件内容（前20字符）: This is the content
```
**提示**: 使用 `os.ReadFile` 一次性读取，或用 `os.Open` + `bufio.Reader`
**参考解法路径**: `data, err := os.ReadFile("filename.txt")`

---

### 6. 写入文件
**目标**: 掌握文件写入操作
**输入输出示例**:
```
写入成功，内容长度: 13
```
**提示**: 使用 `os.WriteFile` 写入，或用 `os.Open` + `Write`
**参考解法路径**: `err := os.WriteFile("output.txt", []byte("Hello"), 0644)`

---

### 7. HTTP Handler 基础
**目标**: 掌握 net/http 的 Handler 写法
**输入输出示例**:
```
启动服务器后访问 http://localhost:8080/hello
响应: Hello, World!
```
**提示**: 实现 `http.Handler` 接口，或用 `http.HandlerFunc`
**参考解法路径**:
```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
})
```

---

### 8. HTTP JSON API
**目标**: 掌握 JSON 请求解析和响应
**输入输出示例**:
```
POST /sum {"a":3,"b":5}
响应: {"result":8}
```
**提示**: 用 `json.NewDecoder(r.Body)` 解析请求，用 `json.NewEncoder(w)` 返回响应
**参考解法路径**:
```go
type Request struct { A, B int }
// 解析请求， 计算， 返回 JSON
```

---

### 9. 错误包装与判断
**目标**: 掌握 errors.Is 和 errors.As
**输入输出示例**:
```
原始错误: permission denied
包装后: operation failed: permission denied
错误包含: true
```
**提示**: `fmt.Errorf` 用 `%w` 包装错误，`errors.Is` 判断错误链
**参考解法路径**: `fmt.Errorf("operation failed: %w", err)`，`errors.Is(err, target)`

---

### 10. 自定义错误类型
**目标**: 创建自定义错误类型
**输入输出示例**:
```
错误类型: *ValidationError
消息: field 'age' must be positive
```
**提示**: 定义一个 struct 实现 `Error()` 方法，实现 `Is` 方法支持 `errors.Is`
**参考解法路径**:
```go
type ValidationError struct {
    Field, Message string
}
func (e *ValidationError) Error() string {
    return fmt.Sprintf("field '%s' %s", e.Field, e.Message)
}
```

---

### 11. 表驱动测试
**目标**: 掌握 Go 推荐的表驱动测试写法
**输入输出示例**:
```
=== RUN   TestAdd
   TestAdd: ok
=== RUN   TestSubtract
    TestSubtract: ok
```
**提示**: 用切片存放测试用例，用 for 循环执行，表驱动测试是 Go 的惯用模式
**参考解法路径**:
```go
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
        got := add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("got %v, want %v", got, tt.want)
        }
    })
}
```

---

### 12. 用 httptest 测试 HTTP
**目标**: 使用 httptest 进行 HTTP 测试
**输入输出示例**:
```
GET /hello 响应状态: 200 OK
响应体: Hello, test!
```
**提示**: 用 `httptest.NewRequest` 创建请求，用 `httptest.NewRecorder` 捕获响应
**参考解法路径**:
```go
req := httptest.NewRequest("GET", "/hello", nil)
w := httptest.NewRecorder()
handler(w, req)
fmt.Println(w.Body.String())
```

---

### 13. goroutine 基础
**目标**: 掌握 goroutine 的创建和启动
**输入输出示例**:
```
启动 3 个 goroutine
goroutine 1 完成
goroutine 2 完成
goroutine 3 完成
```
**提示**: 用 `go` 关键字启动 goroutine
**参考解法路径**: `go func() { fmt.Println("goroutine 完成") }()`

---

### 14. channel 发送接收
**目标**: 掌握 channel 的基本操作
**输入输出示例**:
```
发送: 42
接收: 42
```
**提示**: 用 `make(chan int)` 创建 channel，用 `<-` 发送和接收
**参考解法路径**: `ch := make(chan int); ch <- 42; v := <-ch`

---

### 15. buffered channel
**目标**: 掌握带缓冲区的 channel
**输入输出示例**:
```
缓冲队列长度: 3
已放入: 3
取出: 1
```
**提示**: `make(chan int, 3)` 创建带缓冲的 channel，不会阻塞直到缓冲区满
**参考解法路径**: `ch := make(chan int, 3)`

---

### 16. select 多路复用
**目标**: 掌握 select 处理多个 channel
**输入输出示例**:
```
收到 ch1 的值: 100
```
**提示**: select 类似 switch，但 case 是 channel 操作
**参考解法路径**:
```go
select {
case v := <-ch1:
    fmt.Println("收到 ch1 的值:", v)
case ch2 <- 200:
    fmt.Println("发送到 ch2")
}
```

---

### 17. Worker Pool
**目标**: 实现 worker pool 模式
**输入输出示例**:
```
任务数: 10, Worker 数: 3
Worker 1 处理: 任务 1
Worker 2 处理: 任务 2
Worker 3 处理: 任务 3
...
所有任务完成
```
**提示**: 启动多个 worker 从 channel 读取任务，任务完成后写入结果 channel
**参考解法路径**: 创建一个 `jobs` channel 和一个 `results` channel，启动 N 个 worker

---

### 18. fan-out/fan-in
**目标**: 掌握并发模式中的 fan-out/fan-in
**输入输出示例**:
```
处理完成，结果: [2 4 6 8 10]
```
**提示**: fan-out 是多个 goroutine 从同一个 channel 读取，fan-in 是合并多个 channel
**参考解法路径**: 启动多个 worker 从输入 channel 读取，写入输出 channel

---

### 19. sync.WaitGroup
**目标**: 掌握 WaitGroup 等待多个 goroutine
**输入输出示例**:
```
启动 5 个任务
等待完成...
所有任务完成
```
**提示**: `wg.Add(1)` 增加计数，`wg.Done()` 标记完成，`wg.Wait()` 阻塞等待
**参考解法路径**:
```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        fmt.Println("任务完成")
    }()
}
wg.Wait()
```

---

### 20. sync.Mutex 保护共享资源
**目标**: 掌握互斥锁的使用
**输入输出示例**:
```
Counter: 1000
```
**提示**: 用 `sync.Mutex` 保护对共享变量的访问，防止数据竞争
**参考解法路径**:
```go
var mu sync.Mutex
var counter int
mu.Lock()
counter++
mu.Unlock()
```

---

## 继续学习

完成中级练习后，继续 [advanced/](../advanced/README.md) 进阶练习。