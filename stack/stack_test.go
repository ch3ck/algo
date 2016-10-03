package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var data = []int{0, 1, 2, 3, 4}
	s := new(Stack)

	for i := 0; i < len(data); i++ {
		s.Push(data[i])
	}

	N := s.Len()
	if N != 5 {
		t.Error("stack length should be 5")
	}

	for i := 0; i < N; i++ {
		s.Pop()
	}

	if !s.IsEmpty() {
		t.Error("stack should be empty")
	}

	for i := 0; i < len(data)-1; i++ {
		s.Push(data[i])
	}

	if s.IsEmpty() {
		t.Error("stack should not be empty")
	}

	if s.Peek() != 3 {
		t.Error("peek should return 3")
	}

	s.Init()

	if !s.IsEmpty() {
		t.Error("stack should be empty")
	}
}
