package main

import (
	"context"
	"fmt"
	"time"
)

func Watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s 退出 ===\n", name)
			return
		default:
			fmt.Printf("%s 正在监听...\n", name)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(4*time.Second),
	)
	defer cancel()
	go Watch(ctx, "go1")
	go Watch(ctx, "go2")
	time.Sleep(6 * time.Second)
	fmt.Println("结束了!!")
}
