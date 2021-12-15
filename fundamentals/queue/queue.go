package queue

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Size() int
	IsEmpty() bool
}

// We use linked list as internal data structure
// to get O(1) speed for push and pop opeartions
type node[Item any] struct {
	item Item
	next *node[Item]
}

type queue[OfType any] struct {
	size int
	head *node[OfType]
	tail *node[OfType]
}

func NewQueue[OfType any]() Queue[OfType] {
	return &queue[OfType]{}
}

func (q *queue[T]) Enqueue(item T) {
	oldTail := q.tail
	q.tail = &node[T]{item: item, next: nil}
	if q.IsEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.size++
}

func (q *queue[T]) Dequeue() T {
	first := q.head
	q.head = q.head.next
	if q.IsEmpty() {
		q.tail = nil
	}
	q.size--
	return first.item
}

func (q *queue[T]) Size() int {
	return q.size
}

func (q *queue[T]) IsEmpty() bool {
	return q.head == nil
}
