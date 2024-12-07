package helpers

import "fmt"

type Queue[T any] struct {
	items []T
	head  int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
	fmt.Printf("New queuelength: %d\n", len(q.items))
}

func (q *Queue[T]) Dequeue() T {
	value := q.items[q.head]
	q.head++
	if q.head > len(q.items)/2 {
		q.items = q.items[q.head:]
		q.head = 0
	}
	return value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head >= len(q.items)
}
