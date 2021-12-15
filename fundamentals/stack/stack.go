package stack

type Stack[Item any] interface {
	Push(Item)
	Pop() Item
	Size() int
	IsEmpty() bool
}

// We use linked list as internal data structure
// to get O(1) speed for push and pop opeartions
type node[Item any] struct {
	item Item
	next *node[Item]
}

type stack[OfType any] struct {
	size int
	head *node[OfType]
}

func NewStack[OfType any]() Stack[OfType] {
	return &stack[OfType]{}
}

func (s *stack[Item]) Push(item Item) {
	next := s.head
	s.head = &node[Item]{item, next}
	s.size++
}

func (s *stack[Item]) Pop() Item {
	head := s.head
	s.head = head.next
	s.size--
	return head.item
}

func (s *stack[_]) Size() int {
	return s.size
}

func (s *stack[_]) IsEmpty() bool {
	return s.size == 0
}
