package main

import "fmt"

var gStr string
var gInt int

func main() {
	var lStr string
	var lInt int
	lStr = "hello"
	lInt = 3

	// 数组初始化
	var strArr = [10]string{"aa", "bb", "cc"}
	// 切片初始化
	var sliceArr = make([]string, 0)
	// 字典初始化
	var dic = map[string]int{
		"apple":      1,
		"watermelon": 2,
	}
	fmt.Println(gStr, gInt, lStr, lInt)
	fmt.Println(strArr, sliceArr, dic)
}
