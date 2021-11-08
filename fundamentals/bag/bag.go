package bag

// "Generic" item.
// When type parameters will land this type became
// type Item[T any] T
type Item interface{}

type Bag interface {
	Add(Item)
	Size() int
	IsEmpty() bool
	ForEach(func(Item))
}

// We use linked list as internal data structure
// to get O(1) speed for Add operation
type node struct {
	item Item
	next *node
}

type bag struct {
	size int
	head *node
}

func NewBag() Bag {
	return &bag{}
}

func (b *bag) Add(item Item) {
	next := b.head
	b.head = &node{item, next}
	b.size++
}

func (b *bag) Size() int {
	return b.size
}

func (b *bag) IsEmpty() bool {
	return b.size == 0
}

// As for now Go doesn't have iterators.
// But we can simulate them with ForEach method
func (b *bag) ForEach(f func(Item)) {
	for current := b.head; current != nil; current = current.next {
		f(current.item)
	}
}
