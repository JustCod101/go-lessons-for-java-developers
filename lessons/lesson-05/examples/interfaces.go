package main

import "fmt"

// Speaker 接口定义行为
type Speaker interface {
	Speak() string
}

// Person 结构体定义数据
type Person struct {
	Name string
}

// Speak 实现 Speaker 接口 (隐式实现)
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Dog 结构体
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func SaySomething(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	p := Person{Name: "Gopher"}
	d := Dog{}

	SaySomething(p)
	SaySomething(d)
}
