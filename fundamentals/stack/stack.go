package stack

// "Generic" item.
// When type parameters will land this type became
// type Item[T any] T
type Item interface{}

type Stack interface {
	Push(Item)
	Pop() Item
	Size() int
	IsEmpty() bool
}

// We use linked list as internal data structure
// to get O(1) speed for push and pop opeartions
type node struct {
	item Item
	next *node
}

type stack struct {
	size int
	head *node
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Push(item Item) {
	next := s.head
	s.head = &node{item, next}
	s.size++
}

func (s *stack) Pop() Item {
	head := s.head
	s.head = head.next
	s.size--
	return head.item
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}
