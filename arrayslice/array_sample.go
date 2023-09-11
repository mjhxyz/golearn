package main

import "fmt"

func main() {
	var arr1 = [...]int{1, 2, 3, 4} // ... 默认为数组长度
	fmt.Println(len(arr1))          // 4
	// arr1[5] // panic: runtime error: index out of range [5] with length 4
	fmt.Println(arr1)
	var arr2 = [10]int{1, 2, 3, 4}
	fmt.Println(arr2) // [1 2 3 4 0 0 0 0 0 0]

	slice1 := arr1[0:2] // [1 2]
	fmt.Println(slice1)
}
