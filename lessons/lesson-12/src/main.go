package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware 示例
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", UserHandler)

	// 包装中间件
	handler := loggingMiddleware(mux)

	fmt.Println("Server starting at :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
