package main

import (
	"fmt"
	_ "golearn/init_sample/pkg1"
)

func init() {
	fmt.Println("main init 1")
}

func init() {
	fmt.Println("main init 2")
}

func main() {
	fmt.Println("main main")
}
