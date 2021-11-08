package queue

// "Generic" item.
// When type parameters will land this type became
// type Item[T any] T
type Item interface{}

type Queue interface {
	Enqueue(Item)
	Dequeue() Item
	Size() int
	IsEmpty() bool
}

// We use linked list as internal data structure
// to get O(1) speed for push and pop opeartions
type node struct {
	item Item
	next *node
}

type queue struct {
	size int
	head *node
	tail *node
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) Enqueue(item Item) {
	oldTail := q.tail
	q.tail = &node{item: item, next: nil}
	if q.IsEmpty() {
		q.head = q.tail
	} else {
		oldTail.next = q.tail
	}
	q.size++
}

func (q *queue) Dequeue() Item {
	first := q.head
	q.head = q.head.next
	if q.IsEmpty() {
		q.tail = nil
	}
	q.size--
	return first.item
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.head == nil
}
