package main

import (
	"errors"
	"fmt"
)

// Divide 演示多返回值和错误处理
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	_, err = Divide(10, 0)
	if err != nil {
		fmt.Println("Expected Error:", err)
	}
}
