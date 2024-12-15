package helpers

import (
	"container/heap"
)

// Item represents a single element in the priority queue.
type Item[T any] struct {
	Value    T   // The value of the item; can be any type.
	Priority int // The priority of the item; lower values indicate higher priority.
	Index    int // The index of the item in the heap, maintained by the heap.Interface.
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*Item[T]

// Len returns the number of elements in the priority queue.
func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

// Less compares two items based on their priority. Lower priority values are considered smaller.
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

// Swap swaps the elements at indexes i and j.
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an element to the heap. Push is called by heap.Push.
func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.Index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the smallest element from the heap. Pop is called by heap.Pop.
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // Avoid memory leak.
	item.Index = -1 // For safety.
	*pq = old[0 : n-1]
	return item
}

// Update modifies the priority and value of an item in the heap.
func (pq *PriorityQueue[T]) Update(item *Item[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// MinPriorityQueue wraps the PriorityQueue to provide a clean interface.
type MinPriorityQueue[T any] struct {
	pq PriorityQueue[T]
}

// NewMinPriorityQueue creates a new MinPriorityQueue.
func NewMinPriorityQueue[T any]() *MinPriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return &MinPriorityQueue[T]{pq: *pq}
}

// Insert adds a value with a given priority to the queue.
func (mpq *MinPriorityQueue[T]) Insert(value T, priority int) {
	item := &Item[T]{
		Value:    value,
		Priority: priority,
	}
	heap.Push(&mpq.pq, item)
}

// ExtractMin removes and returns the element with the lowest priority.
func (mpq *MinPriorityQueue[T]) ExtractMin() (T, int, bool) {
	//if len(mpq.pq) == 0 {
	//	return mpq.pq[0], 0, false
	//}
	item := heap.Pop(&mpq.pq).(*Item[T])
	return item.Value, item.Priority, true
}

// IsEmpty checks if the priority queue is empty.
func (mpq *MinPriorityQueue[T]) IsEmpty() bool {
	return len(mpq.pq) == 0
}
