package stack

import (
	"sync"
)

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (s *ArrayStack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.array = append(s.array, v)
	s.size += 1
}

func (s *ArrayStack) Pop() string {
	if s.size == 0 {
		panic("stack is empty")
	}
	v := s.array[s.size-1]
	newArray := make([]string, s.size-1)
	for i := 0; i < s.size-1; i++ {
		newArray[i] = s.array[i]
	}
	s.array = newArray
	// 栈中元素 -1
	s.size--
	return v
}

func (s *ArrayStack) Peek() string {
	if s.size == 0 {
		panic("stack is empty")
	}

	v := s.array[s.size-1]
	return v
}

func (s *ArrayStack) IsEmpty() bool {
	return s.size == 0
}

func (s *ArrayStack) Size() int {
	return s.size
}
