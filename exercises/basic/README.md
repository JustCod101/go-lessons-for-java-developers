# 基础练习

本目录包含 Go 语言基础练习，适合刚学完前 5 节课程的学习者。

---

## 练习清单

### 1. Hello World
**目标**: 验证 Go 环境是否正常工作
**输入输出示例**:
```
$ go run main.go
Hello, Go!
```
**提示**: 最简单的练习，看看你的 Go 环境是否正常
**参考解法路径**: 创建 `main.go`，使用 `fmt.Println("Hello, Go!")`

---

### 2. 变量声明与类型推断
**目标**: 掌握 Go 的变量声明方式
**输入输出示例**:
```
输入: 无
输出: name=Alice age=30 score=95.5
```
**提示**: 使用 `var`、`:=` 两种方式，感受类型推断
**参考解法路径**: `var name string = "Alice"; age := 30; var score float64 = 95.5`

---

### 3. 常量与 iota
**目标**: 理解 Go 的常量定义和 iota 枚举
**输入输出示例**:
```
Monday=1
Tuesday=2
Wednesday=3
```
**提示**: iota 在 const 块内从 0 开始递增
**参考解法路径**: `const (Monday = iota + 1; Tuesday; Wednesday)`

---

### 4. 基本数据类型转换
**目标**: 掌握 Go 的类型转换语法
**输入输出示例**:
```
int to float: 42.0
float to int: 10
int to string: 123
```
**提示**: Go 没有隐式转换，必须显式调用类型转换函数
**参考解法路径**: `float64(42)`, `int(10.9)`, `strconv.Itoa(123)`

---

### 5. 字符串操作
**目标**: 熟悉 strings 包常用函数
**输入输出示例**:
```
原始: Hello World
转大写: HELLO WORLD
转小写: hello world
包含 "Go": true
分割: [Hello World]
```
**提示**: strings 包提供了 ToUpper、ToLower、Contains、Split
**参考解法路径**: `strings.ToUpper(s)`, `strings.Contains(s, "Go")`

---

### 6. if 条件判断
**目标**: 掌握 Go 的 if 语句语法
**输入输出示例**:
```
输入: age = 20
输出: 成年人
```
**提示**: Go 的 if 条件不加括号，可以包含初始化语句
**参考解法路径**:
```go
if age := 20; age >= 18 {
    fmt.Println("成年人")
}
```

---

### 7. for 循环的三种形式
**目标**: 掌握 Go 唯一的循环关键字 for 的用法
**输入输出示例**:
```
标准循环: 0 1 2 3 4
条件循环: 0 1 2 3
无限循环（只取前5个）: 0 1 2 3 4
```
**提示**: Go 没有 while，用 for 代替；for 的三种形式
**参考解法路径**: `for i:=0; i<5; i++`, `for i<5`, `for {}`

---

### 8. switch 多分支
**目标**: 掌握 Go 的 switch 语句
**输入输出示例**:
```
grade=A
```
**提示**: switch 不加 break 也会自动 break，可以写 fallthrough 穿透
**参考解法路径**:
```go
switch score := 85; {
case score >= 90: fmt.Println("A")
case score >= 80: fmt.Println("B")
default: fmt.Println("C")
}
```

---

### 9. defer 延迟执行
**目标**: 理解 defer 的执行时机
**输入输出示例**:
```
start
end
deferred
```
**提示**: defer 在函数 return 前执行，多个 defer 按栈顺序执行
**参考解法路径**:
```go
defer fmt.Println("deferred")
fmt.Println("start")
fmt.Println("end")
```

---

### 10. 指针基础
**目标**: 理解 Go 的指针用法
**输入输出示例**:
```
原始值: 10
通过指针修改后: 20
```
**提示**: `*` 表示指针类型，`&` 获取变量地址
**参考解法路径**:
```go
x := 10
p := &x
*p = 20
```

---

### 11. 函数定义与调用
**目标**: 掌握 Go 函数的定义方式
**输入输出示例**:
```
add(3, 5) = 8
```
**提示**: 函数可以有多个参数和多个返回值
**参考解法路径**: `func add(a, b int) int { return a + b }`

---

### 12. 多返回值函数
**目标**: 掌握 Go 的多返回值特性
**输入输出示例**:
```
商=3 余=1
```
**提示**: 函数可以返回多个值，用 `,` 分隔
**参考解法路径**: `func div(a, b int) (int, int) { return a / b, a % b }`

---

### 13. 命名返回值
**目标**: 理解 Go 的命名返回值
**输入输出示例**:
```
sum=15
```
**提示**: 命名返回值会作为已声明变量在函数内使用
**参考解法路径**:
```go
func sum(a, b, c int) (result int) {
    result = a + b + c
    return // 直接 return result
}
```

---

### 14. 匿名函数与闭包
**目标**: 理解匿名函数和闭包的概念
**输入输出示例**:
```
闭包计算: 3 + 5 = 8
```
**提示**: 匿名函数可以赋值给变量，闭包可以访问外部变量
**参考解法路径**:
```go
add := func(a, b int) int { return a + b }
fmt.Println(add(3, 5))
```

---

### 15. error 处理
**目标**: 掌握 Go 的错误处理模式
**输入输出示例**:
```
错误: open nonexistent.txt: no such file or directory
```
**提示**: Go 的 error 是一个接口，实现 `Error() string` 方法即可
**参考解法路径**:
```go
f, err := os.Open("nonexistent.txt")
if err != nil {
    fmt.Println("错误:", err)
}
```

---

### 16. struct 定义与使用
**目标**: 掌握 Go 的 struct 定义
**输入输出示例**:
```
Name: Alice, Age: 30
```
**提示**: Go 没有 class，用 struct 代替；用工厂函数代替构造函数
**参考解法路径**:
```go
type Person struct {
    Name string
    Age  int
}
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}
```

---

### 17. struct 方法（value receiver）
**目标**: 掌握为 struct 定义方法
**输入输出示例**:
```
原始: {Alice 30}
字符串: Person{Name:Alice, Age:30}
```
**提示**: 方法的 receiver 如果是 value 类型，方法内无法修改原始 struct
**参考解法路径**:
```go
func (p Person) String() string {
    return fmt.Sprintf("Person{Name:%s, Age:%d}", p.Name, p.Age)
}
```

---

### 18. struct 方法（pointer receiver）
**目标**: 掌握用 pointer receiver 修改 struct
**输入输出示例**:
```
修改前: Alice age=30
修改后: Alice age=31
```
**提示**: 如果方法需要修改 struct，必须用 pointer receiver
**参考解法路径**:
```go
func (p *Person) Birthday() {
    p.Age++
}
```

---

### 19. 隐式接口实现
**目标**: 理解 Go 的接口隐式实现
**输入输出示例**:
```
实现了 Stringer 接口: true
```
**提示**: Go 的接口不需要显式声明实现，只要方法签名匹配即可
**参考解法路径**:
```go
type Stringer interface {
    String() string
}
type Person struct{}
func (p Person) String() string { return "Person" }
var _ Stringer = Person{}
```

---

### 20. 空接口（any）
**目标**: 理解空接口和类型断言
**输入输出示例**:
```
空接口可以存储: 42 (type: int)
空接口可以存储: hello (type: string)
```
**提示**: 空接口 `any` 可以存储任何类型，用类型断言获取具体值
**参考解法路径**:
```go
var i any = 42
fmt.Printf("空接口可以存储: %v (type: %T)\n", i, i)
```

---

## 继续学习

完成基础练习后，继续 [intermediate/](../intermediate/README.md) 中级练习。