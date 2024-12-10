package helpers

type Stack[T any] struct {
	items []StackItem[T]
}

func NewStack[T any]() Stack[T] {
	return Stack[T]{items: make([]StackItem[T], 0)}
}

func (s *Stack[T]) Push(value StackItem[T]) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() StackItem[T] {
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Peek() StackItem[T] {
	return s.items[len(s.items)-1]
}

type StackItem[T any] struct {
	Accumulator      T
	Trace            []T
	Value            T
	TraceAccumulator []T
}

func (si StackItem[T]) GetChild(operation func(T, T) T, newValue T) StackItem[T] {
	newAccumulator := operation(si.Accumulator, newValue)
	return StackItem[T]{
		Accumulator:      newAccumulator,
		Value:            newValue,
		Trace:            append(si.Trace, newValue),
		TraceAccumulator: append(si.TraceAccumulator, newAccumulator),
	}
}
