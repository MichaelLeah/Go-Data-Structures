// Package godatastructures provides basic implementation of various data structures such as Stacks and Queues.
package godatastructures

// Stack offers a last in, first out storage mechanism.
type Stack struct {
	data       []interface{} // Stores any data type
	topOfStack int           // Last entry in the stack
}

// Push adds n items onto the stack in the order they are given
func (s *Stack) Push(items ...interface{}) {
	for i := range items {
		s.data = append(s.data, items[i])
	}
	s.topOfStack = (len(s.data) - 1)
}

// Pop removes the last item added (the top of the stack) from the stack
func (s *Stack) Pop() {
	s.data = s.data[:s.topOfStack]
	s.topOfStack = (len(s.data) - 1)
}

// Peek returns the last item added (the top of the stack) from the stack
func (s *Stack) Peek() interface{} {
	return s.data[s.topOfStack]
}
