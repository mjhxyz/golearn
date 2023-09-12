package queue

import (
	"sync"
)

type Queue struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (q *Queue) Put(v string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.array = append(q.array, v)
	q.size += 1
}

func (q *Queue) Poll() (v string) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.size == 0 {
		panic("queue is empty")
	}

	v = q.array[0]
	q.array = q.array[1:]
	q.size--
	return
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
