package queuestack

import "testing"

func TestStackIsInitiallyEmpty(t *testing.T) {
	s := NewStack()

	if !s.Empty() {
		t.Error("newly created stack should contain no elements!")
	}
}

func TestStackIsNonEmptyAfterPush(t *testing.T) {
	s := NewStack()

	s.Push(1)

	if s.Empty() {
		t.Error("stack should not be empty after pushing any elements!")
	}
}

func TestStackIsEmptyAfterPushAndPop(t *testing.T) {
	s := NewStack()

	s.Push(1)
	v, err := s.Pop()

	if !s.Empty() {
		t.Error("stack should be empty after pushing and popping an element!")
	}

	if err != nil {
		t.Error("there should not have been anything erroneous occurring after pushing and popping an element:", err)
	}

	if v != 1 {
		t.Error("value popped should be the same as what was pushed! got:", v, "want:", 1)
	}
}

func TestPoppingAnEmptyStackErrors(t *testing.T) {
	s := NewStack()

	_, err := s.Pop()

	if err == nil {
		t.Error("stack should have errored when it was popped while empty!")
	}
}

func TestPeekingAnEmptyStackErrors(t *testing.T) {
	s := NewStack()

	_, err := s.Peek()

	if err == nil {
		t.Error("stack should have errored when it was peeked while empty!")
	}
}

func TestPeekReturnsTopOfStack(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push(2)

	v, err := s.Peek()

	if err != nil {
		t.Error("there should have been nothing erroneous that occurred:", err)
	}

	if v != 2 {
		t.Error("stack should have returned the top value on the stack! got:", v, "want:", 2)
	}
}

func TestStackIsLIFOCompliant(t *testing.T) {
	s := NewStack()

	values := []interface{}{0, 1, 2, 3}

	for _, v := range values {
		s.Push(v)
	}

	offset := len(values) - 1
	for !s.Empty() {
		val, err := s.Pop()

		if err != nil {
			t.Fatal("encountered unexpected error:", err)
		}

		if val != values[offset] {
			t.Error("stack is not LIFO compliant")
		}

		offset--
	}
}
