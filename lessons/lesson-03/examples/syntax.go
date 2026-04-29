package main

import "fmt"

func main() {
	// 变量声明
	var name string = "Java Developer"
	age := 25 // 简短声明

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 循环 (Go 只有 for)
	fmt.Print("Counting: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 指针
	x := 10
	p := &x
	fmt.Printf("Value of x: %d, Address: %p, Value via pointer: %d\n", x, p, *p)
}
