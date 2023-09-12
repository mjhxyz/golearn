package main

import (
	"fmt"
	"golearn/stack_sample/queue"
	"golearn/stack_sample/stack"
)

func main() {
	arrayStack := new(stack.ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("============================")

	q := new(queue.Queue)
	q.Put("1")
	q.Put("2")
	q.Put("3")
	q.Put("4")
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
	q.Put("5")
	fmt.Println(q.Poll())
	fmt.Println(q.Poll())
}
