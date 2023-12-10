package main

import "fmt"

func main() {
	lStr := "case3"
	if lStr == "case3" {
		fmt.Println("就是 case3")
	} else {
		fmt.Println("不是 case3")
	}

	// 字典初始化
	var dic = map[string]int{
		"j": 1,
		"k": 2,
	}
	if num, ok := dic["j"]; ok {
		fmt.Printf("j是%d\n", num)
	}

	switch lStr {
	case "case1":
		fmt.Println("是 case1")
	case "case2":
		fmt.Println("是 case2")
	case "case3":
		fmt.Println("是 case3")
	}
}
