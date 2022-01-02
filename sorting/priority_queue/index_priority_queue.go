package priority_queue

// TODO: change name to "keyed priority queue"
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
	qp   []int
	less func(T, T) bool
}

// TODO: panic for illegal operations
// TODO: can we construct queue without bounded index size?
func NewIPQ[T any](less func(T, T) bool, indexSize int) IndexPriorityQueue[T] {
	n := 0
	// TODO: switch to 0 based index
	items := make([]T, indexSize+1)
	pq := make([]int, indexSize+1)
	qp := make([]int, indexSize+1)

	for i := range qp {
		qp[i] = -1
	}

	for i := range pq {
		pq[i] = -1
	}

	return &indexPriorityQueue[T]{n, items, pq, qp, less}
}

func (q *indexPriorityQueue[T]) top() T {
	return q.items[q.pq[1]]
}

func (q *indexPriorityQueue[T]) topKey() int {
	return q.pq[1]
}

func (q *indexPriorityQueue[T]) insert(key int, item T) {
	q.n++
	q.pq[q.n] = key
	q.qp[key] = q.n
	q.items[key] = item
	q.swim(q.n)
}

func (q *indexPriorityQueue[T]) remove() T {
	topKey := q.topKey()
	q.swap(1, q.n)
	q.n--
	q.sink(1)
	// TODO: need to remove actual item from items array
	// to allow removed item to be GCed
	q.qp[q.pq[q.n+1]] = -1
	return q.items[topKey]
}

func (q *indexPriorityQueue[T]) removeKey(key int) T {
	pivot := q.qp[key]
	q.swap(pivot, q.n)
	q.n--
	q.swim(pivot)
	q.sink(pivot)
	// TODO: need to remove actual item from items array
	// to prevent memory leak
	q.qp[key] = -1
	q.pq[q.n+1] = -1
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
	for 2*parent <= q.n {
		child := 2 * parent
		if child < q.n && q.less(q.items[q.pq[child]], q.items[q.pq[child+1]]) {
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
	for child > 1 && q.less(q.items[q.pq[child/2]], q.items[q.pq[child]]) {
		q.swap(child/2, child)
		child = child / 2
	}
}

func (q *indexPriorityQueue[_]) swap(i, j int) {
	q.qp[q.pq[i]] = j
	q.qp[q.pq[j]] = i
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
}
