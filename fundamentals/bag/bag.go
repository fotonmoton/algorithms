package bag

type Bag[Item any] interface {
	Add(Item)
	Size() int
	IsEmpty() bool
	ForEach(func(Item))
}

// We use linked list as internal data structure
// to get O(1) speed for Add operation
type node[Item any] struct {
	item Item
	next *node[Item]
}

type bag[OfType any] struct {
	size int
	head *node[OfType]
}

func NewBag[OfType any]() Bag[OfType] {
	return &bag[OfType]{}
}

func (b *bag[OfType]) Add(item OfType) {
	next := b.head
	b.head = &node[OfType]{item, next}
	b.size++
}

func (b *bag[_]) Size() int {
	return b.size
}

func (b *bag[_]) IsEmpty() bool {
	return b.size == 0
}

// As for now Go doesn't have iterators.
// But we can simulate them with ForEach method
func (b *bag[OfType]) ForEach(f func(OfType)) {
	for current := b.head; current != nil; current = current.next {
		f(current.item)
	}
}
