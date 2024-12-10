package helpers

type Queue[T any] struct {
	items []QueueItem[T]
	head  int
}

type QueueItem[T any] struct {
	Value  T
	Traces []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]QueueItem[T], 0)}
}

func (q *Queue[T]) Enqueue(value QueueItem[T]) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Dequeue() QueueItem[T] {
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

func (q *Queue[T]) Peek() QueueItem[T] {
	return q.items[q.head]
}
