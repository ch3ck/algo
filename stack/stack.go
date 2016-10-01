package stack

// Element is an element of a stack
type Element struct {
	// The value stored with this element
	Value interface{}
}

// Stack represents a stack
type Stack struct {
	items  []Element // items of your stack
	length int       // current stack length
}

// Init initializes or clears the stack s
func (s *Stack) Init() *Stack {
	s.items = nil
	s.length = 0
	return s
}

// New returns an initialized stack
func New() *Stack {
	return new(Stack).Init()
}

// Length returns the number of elements of the stack s
// The complexity is O(1)
func (s *Stack) Len() int {
	return s.length
}

// IsEmpty returns true if the stack s is empty
func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// Push inserts an item into the stack
func (s *Stack) Push(v interface{}) {
	s.items = append(s.items[:s.length], Element{Value: v})
	s.length++
}

// Pop removes an item from the stack
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	ToRemove := s.items[s.length-1]
	s.items = s.items[:(s.length - 1)]
	s.length--
	return ToRemove.Value
}

// Peek shows you the last element of the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.items[s.length-1].Value
}
