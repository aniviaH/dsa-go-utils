package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	// 使用for循环遍历切片
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// list := NewLinkedListFromSlice(slice)
}
