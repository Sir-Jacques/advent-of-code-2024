package helpers

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() T {
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

type StackItem[T any] struct {
	Accumulator      T
	RemainingNumbers []T
}

func (si *StackItem[T]) GetChild(operation func(T, T) T) *StackItem[T] {
	return &StackItem[T]{
		Accumulator:      operation(si.Accumulator, si.RemainingNumbers[0]),
		RemainingNumbers: si.RemainingNumbers[1:]}
}
