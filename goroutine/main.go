package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	f func() error
}

type Pool struct {
	RunningWorkerss int64 // 正在运行的协程数量
	Capacity        int64 // 最大协程数量
	JobChan         chan *Task
	sync.Mutex
}

func (p *Pool) AddTask(task *Task) {
	// 需要保证原子性: 两个协程同时调用 AddTask
	// A 判断协程数量后，让出时间片
	// B 继续判断也能成功
	// 如果 A 的时候，已经刚好达到最大，那么 B 的话，就会超出设置的最大协程数量
	p.Lock()
	defer p.Unlock()

	// 如果正在运行的协程数量小于最大协程数量, 就运行一个
	if p.RunningWorkerss < p.Capacity {
		time.Sleep(1 * time.Second) // 模拟让出时间片的情况
		p.Run()                     // 启动一个新的协程处理
	}
	fmt.Printf("当前协程数量: %d\n", p.RunningWorkerss)

	p.JobChan <- task // 如果还有额外的协程，那么直接能执行，否则会阻塞
}

func (p *Pool) Run() {
	atomic.AddInt64(&p.RunningWorkerss, 1)
	go func() {
		defer func() {
			atomic.AddInt64(&p.RunningWorkerss, -1)
		}()
		for task := range p.JobChan {
			task.f()
		}
	}()
}

func NewPool(capacity int64, taskNum int) *Pool {
	return &Pool{
		Capacity: capacity,
		JobChan:  make(chan *Task, taskNum),
	}
}

func NewTask(f func() error) *Task {
	return &Task{
		f: f,
	}
}

func main() {
	pool := NewPool(3, 10)
	for i := 0; i < 20; i++ {
		go func() {
			pool.AddTask(NewTask(func() error {
				fmt.Println("i'm task")
				time.Sleep(5 * time.Second)
				return nil
			}))
		}()
	}
	fmt.Println("创建任务完成...")
	time.Sleep(100 * time.Second)
}
