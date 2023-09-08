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
			fmt.Printf("%s 结束了.\n", name)
			return
		default:
			fmt.Printf("%s 继续监听...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()

	go Watch(ctx, "go1")
	go Watch(ctx, "go2")

	time.Sleep(6 * time.Second)
}
