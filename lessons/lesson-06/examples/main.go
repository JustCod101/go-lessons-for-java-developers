package main

import (
	"fmt"
)

func main() {
	// 1. Array: 固定长度
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("Array: %v, len: %d\n", arr, len(arr))

	// 2. Slice: 动态数组
	// Java 对比: 类似于 ArrayList
	slice := []int{1, 2, 3}
	fmt.Printf("Initial Slice: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	// append 扩容
	slice = append(slice, 4, 5)
	fmt.Printf("After append: %v, len: %d, cap: %d\n", slice, len(slice), cap(slice))

	// 3. Map
	// Java 对比: 类似于 HashMap
	m := make(map[string]int)
	m["apple"] = 1
	m["banana"] = 2
	fmt.Printf("Map: %v\n", m)

	// 检查 key 是否存在
	val, ok := m["orange"]
	if !ok {
		fmt.Println("Key 'orange' not found")
	} else {
		fmt.Println("Orange value:", val)
	}

	// 4. Set (Go 没有内置 Set，通常用 map[K]struct{} 实现)
	set := make(map[string]struct{})
	set["item1"] = struct{}{}
	set["item2"] = struct{}{}

	if _, exists := set["item1"]; exists {
		fmt.Println("Item1 exists in set")
	}

	// 5. Range 遍历
	fmt.Println("Iterating over slice:")
	for i, v := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
