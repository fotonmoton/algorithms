package priority_queue

type IndexPriorityQueue[T any] interface {
	top() T                 // get item with biggest priority
	topKey() int            // get key of an item with biggest priority
	remove() T              // removes item with the biggest priority
	removeKey(key int) T    // removes item with specified key
	insert(key int, item T) // adds item with specified key
	change(key int, item T) // changes item with specified key
	contains(key int) bool  // checks if key exists in queue
	isEmpty() bool
	size() int
}

type indexPriorityQueue[T any] struct {
	n int
	// unordered items.
	// items[key] = item
	items []T
	// priority queue. Contains keys for items in priority order.
	// items[pq[1]] = item with biggest priority
	pq []int
	// "reverse" for pq. Maps item key to priority
	// qp[key] = priority index, qp[pq[key]] = pq[qp[key]] = key
	qp []int
	// "less" function. Depending on desired order first element
	// not necessary less than a second
	less func(T, T) bool
}

// TODO: panic for illegal operations
// TODO: can we construct queue without bounded index size?
func NewIPQ[T any](less func(T, T) bool, indexSize int) IndexPriorityQueue[T] {
	n := 0
	items := make([]T, indexSize)
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

func (q *indexPriorityQueue[T]) top() T {
	return q.items[q.pq[0]]
}

func (q *indexPriorityQueue[T]) topKey() int {
	return q.pq[0]
}

func (q *indexPriorityQueue[T]) insert(key int, item T) {
	q.pq[q.n] = key
	q.qp[key] = q.n
	q.items[key] = item
	q.swim(q.n)
	q.n++
}

func (q *indexPriorityQueue[T]) remove() T {
	return q.removeKey(q.topKey())
}

func (q *indexPriorityQueue[T]) removeKey(key int) T {
	pivot := q.qp[key]
	q.n--
	q.swap(pivot, q.n)
	q.swim(pivot)
	q.sink(pivot)
	// TODO: need to remove actual item from items array
	// to prevent memory leak
	q.qp[key] = -1
	q.pq[q.n] = -1
	return q.items[key]
}

func (q *indexPriorityQueue[T]) change(key int, item T) {
	q.items[key] = item
	q.swim(q.qp[key])
	q.sink(q.qp[key])
}

func (pq *indexPriorityQueue[_]) size() int {
	return pq.n
}

func (pq *indexPriorityQueue[_]) isEmpty() bool {
	return pq.n == 0
}

func (q *indexPriorityQueue[_]) contains(key int) bool {
	return q.qp[key] != -1
}

func (q *indexPriorityQueue[T]) sink(parent int) {

	for {
		child := 2*parent + 1

		if child >= q.n {
			break
		}

		if child+1 < q.n && q.less(q.items[q.pq[child]], q.items[q.pq[child+1]]) {
			child++
		}

		if !q.less(q.items[q.pq[parent]], q.items[q.pq[child]]) {
			break
		}

		q.swap(parent, child)

		parent = child
	}

}

func (q *indexPriorityQueue[T]) swim(child int) {
	for {
		parent := (child - 1) / 2

		if child <= 0 || q.less(q.items[q.pq[child]], q.items[q.pq[parent]]) {
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
