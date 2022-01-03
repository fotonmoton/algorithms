package priority_queue

type IndexPriorityQueue[T any] interface {
	Top() T                    // get item with biggest priority
	TopIndex() int             // get index of an item with biggest priority
	Remove() T                 // removes item with the biggest priority
	RemoveAtIndex(index int) T // removes item at specified index
	Insert(index int, item T)  // adds item at specified index
	Change(index int, item T)  // changes item at specified index and preserves ordering
	Contains(index int) bool   // checks if item exists at specified index
	IsEmpty() bool
	Size() int
}

type indexPriorityQueue[T any] struct {
	n int
	// unordered items.
	// items[index] = item.
	// we store pointers instead of values to prevent memory leaks
	// by setting nil for removed items
	items []*T
	// priority queue. Contains keys for items in priority order.
	// items[pq[1]] = item with biggest priority
	pq []int
	// "reverse" for pq. Maps item index to priority
	// qp[index] = priority index, qp[pq[index]] = pq[qp[index]] = index
	qp []int
	// "less" function. Depending on desired order first element
	// not necessary less than a second
	less func(T, T) bool
}

// TODO: panic for illegal operations
// TODO: can we construct queue without bounded index size?
func NewIPQ[T any](less func(T, T) bool, indexSize int) IndexPriorityQueue[T] {
	n := 0
	items := make([]*T, indexSize)
	pq := make([]int, indexSize)
	qp := make([]int, indexSize)

	for i := range qp {
		qp[i] = -1
	}

	for i := range pq {
		pq[i] = -1
	}

	return &indexPriorityQueue[T]{n, items, pq, qp, less}
}

func (q *indexPriorityQueue[T]) Top() T {
	return *q.items[q.pq[0]]
}

func (q *indexPriorityQueue[T]) TopIndex() int {
	return q.pq[0]
}

func (q *indexPriorityQueue[T]) Insert(index int, item T) {
	q.pq[q.n] = index
	q.qp[index] = q.n
	q.items[index] = &item
	q.swim(q.n)
	q.n++
}

func (q *indexPriorityQueue[T]) Remove() T {
	return q.RemoveAtIndex(q.TopIndex())
}

func (q *indexPriorityQueue[T]) RemoveAtIndex(index int) T {
	pivot := q.qp[index]
	removed := q.items[index]
	q.n--
	q.swap(pivot, q.n)
	q.swim(pivot)
	q.sink(pivot)
	q.items[index] = nil
	q.qp[index] = -1
	q.pq[q.n] = -1
	return *removed
}

func (q *indexPriorityQueue[T]) Change(index int, item T) {
	q.items[index] = &item
	q.swim(q.qp[index])
	q.sink(q.qp[index])
}

func (pq *indexPriorityQueue[_]) Size() int {
	return pq.n
}

func (pq *indexPriorityQueue[_]) IsEmpty() bool {
	return pq.n == 0
}

func (q *indexPriorityQueue[_]) Contains(index int) bool {
	return q.qp[index] != -1
}

func (q *indexPriorityQueue[T]) sink(parent int) {

	for {
		child := 2*parent + 1

		if child >= q.n {
			break
		}

		if child+1 < q.n && q.less(*q.items[q.pq[child]], *q.items[q.pq[child+1]]) {
			child++
		}

		if !q.less(*q.items[q.pq[parent]], *q.items[q.pq[child]]) {
			break
		}

		q.swap(parent, child)

		parent = child
	}

}

func (q *indexPriorityQueue[T]) swim(child int) {
	for {
		parent := (child - 1) / 2

		if child <= 0 || q.less(*q.items[q.pq[child]], *q.items[q.pq[parent]]) {
			break
		}

		q.swap(parent, child)

		child = parent
	}
}

func (q *indexPriorityQueue[_]) swap(i, j int) {
	q.qp[q.pq[i]] = j
	q.qp[q.pq[j]] = i
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
}
