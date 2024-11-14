package condition

// Stack A simple stack implementation for use in the condition resolution algorithm
type Stack[T any] struct {
	data []T
}

// Makes an empty stack
func makeStack[T any]() *Stack[T] {
	return new(Stack[T])
}

// Push to top of stack
func (s *Stack[T]) push(elem T) {
	s.data = append(s.data, elem)
}

// Pop from top of stack, returning the element that was at top
func (s *Stack[T]) pop() T {
	n := len(s.data)
	res := s.data[n-1]
	s.data = s.data[:n-1]
	return res
}
