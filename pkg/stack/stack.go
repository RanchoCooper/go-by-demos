package stack

import (
	"fmt"
	"sync"
)

type Stack struct {
	lock sync.RWMutex
	slice []string
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Push(value string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.slice = append(s.slice, value)
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty")
	}

	s.lock.Lock()
	defer s.lock.Unlock()

	s.slice = s.slice[:s.Size() - 1]
	return nil
}

func (s *Stack) Top() string {
	if s.IsEmpty() {
		return ""
	}
	return s.slice[len(s.slice) - 1]
}
