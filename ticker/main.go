package main

import (
	"fmt"
	"time"
)

func Watcher() chan struct{} {
	ticker := time.NewTicker(2 * time.Second)
	ch := make(chan struct{})

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("invoke ticker")
			case <-ch:
				fmt.Println("stop ticker")
				return
			}
		}
	}()
	return ch
}

func main() {
	done := Watcher()
	time.Sleep(10 * time.Second)
	done <- struct{}{}
	close(done)
}
