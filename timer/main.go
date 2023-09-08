package main

import (
	"fmt"
	"time"
)

func main() {
	// 两秒后超时
	<-time.NewTimer(2 * time.Second).C
	<-time.After(2 * time.Second)
	fmt.Println("2s 后了...")
}
