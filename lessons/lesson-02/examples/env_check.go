package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("OS: %s\n", runtime.GOOS)
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
}
