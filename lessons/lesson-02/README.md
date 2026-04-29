# Lesson 02: 环境与 Go Modules

## 1. 学习目标
- 掌握 Go 开发环境的配置与常用命令
- 理解 Go Modules 的工作原理
- 能够对比 Go Modules 与 Maven/Gradle 的异同
- 掌握 package 与 module 的组织关系

## 2. 给 Java 开发者的类比

### Go Modules vs Maven/Gradle
Java 使用 `pom.xml` 或 `build.gradle` 管理依赖。Go 使用 `go.mod` 记录依赖版本，`go.sum` 记录依赖的哈希值以确保安全。

### Package vs Package
Java 的包名通常与目录结构严格对应。Go 的包名也建议与目录一致，但一个目录下只能有一个包（不含测试包）。Go 的导入路径是基于 module 根路径的。

### Workspace vs GOPATH
早期的 Go 使用 `GOPATH` 模式，类似把所有项目都放在一个全局的 `lib` 下。现在的 Go Modules 模式允许你在任何地方创建项目，类似于 Maven 的本地仓库管理方式。

## 3. 核心概念

### 常用命令
- `go version`: 查看版本。
- `go env`: 查看环境变量（如 `GOPROXY`）。
- `go mod init <name>`: 初始化模块。
- `go run`: 编译并运行。
- `go build`: 编译生成二进制文件。
- `go test`: 运行单元测试。
- `go fmt`: 格式化代码（Go 强制统一代码风格）。

### Go Modules 详解
- **go.mod**: 定义模块路径和依赖列表。
- **go.sum**: 校验文件，防止依赖被篡改。
- **依赖下载**: 默认下载到 `$GOPATH/pkg/mod`，不需要手动管理。

### Java 对比
- **构建速度**: Go 的构建速度远快于 Maven，因为它没有复杂的生命周期插件。
- **依赖冲突**: Go 使用语义化版本控制，处理冲突的方式比 Maven 的“路径最短优先”更直观。

## 4. 代码示例
参考示例文件：`examples/env_check.go`

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Printf("Go Version: %s\n", runtime.Version())
    fmt.Printf("OS: %s\n", runtime.GOOS)
}
```

初始化模块示例：
```bash
go mod init github.com/yourname/myproject
```

## 5. 常见误区
- **误区 1**: 像 Java 一样手动创建复杂的目录结构。Go 提倡扁平化，不要过度设计目录。
- **误区 2**: 忘记设置 `GOPROXY`。在国内开发，不设置代理会导致依赖下载失败。
- **误区 3**: 在一个目录下写多个 package。Go 规定一个文件夹下只能有一个包名。

## 6. 本节练习
1. 在当前目录下执行 `go mod init lesson02`。
2. 运行 `examples/env_check.go` 并查看输出。
3. 尝试使用 `go fmt` 格式化一个故意写乱格式的 Go 文件。

## 7. 面试可能怎么问
- **问**: `go.mod` 和 `go.sum` 的区别是什么？
- **答**: `go.mod` 记录了项目依赖的模块和版本要求；`go.sum` 记录了依赖模块内容的加密哈希，用于验证下载的包是否被篡改。
- **问**: Go 如何处理循环依赖？
- **答**: Go 在编译期严禁包之间的循环依赖。如果出现 A 调 B，B 调 A，编译会直接报错。这强制开发者设计更好的解耦结构。

## 8. 本节总结
掌握 Go Modules 是开始 Go 开发的第一步。告别了 Maven 的繁琐配置，你会发现 Go 的依赖管理非常纯粹。记住：一个目录一个包，一个项目一个 mod。
