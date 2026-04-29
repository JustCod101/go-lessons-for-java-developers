package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "Java Developer"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func main() {
	// 1. fmt: 格式化输入输出
	fmt.Println("Standard Library Demo")

	// 2. time: 时间处理
	// Java 对比: 类似于 java.time
	now := time.Now()
	fmt.Println("Current time:", now.Format("2006-01-02 15:04:05"))

	// 3. os: 操作系统交互
	hostname, _ := os.Hostname()
	fmt.Println("Hostname:", hostname)

	// 4. encoding/json: JSON 处理
	// Java 对比: 类似于 Jackson 或 Gson
	u := User{ID: 100, Name: "Gopher"}
	data, _ := json.Marshal(u)
	fmt.Println("JSON:", string(data))

	// 5. net/http: HTTP 服务器
	// Java 对比: 类似于 Spring Boot (但更底层)
	fmt.Println("Starting server on :8080...")
	http.HandleFunc("/user", handler)

	// 为了演示不阻塞，这里用 goroutine 启动，实际应用中直接调用
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Server is running (simulated)")
}
