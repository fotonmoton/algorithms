package sorting

type PriorityQueue[T any] interface {
	top() T
	insert(T)
	delete() T
	size() int
	isEmpty() bool
}

type priorityQueue[T any] struct {
	n    int
	heap []T
	less func(T, T) bool
}

func NewPQ[T any](less func(T, T) bool) PriorityQueue[T] {
	return &priorityQueue[T]{
		n:    0,
		less: less,
		// First element is not used
		// to make equation "next = current * 2" work
		heap: make([]T, 1, 1),
	}
}

func (pq *priorityQueue[T]) top() T {
	// After all operations biggest element
	// should be always at the beginning of the array
	return pq.heap[1]
}

func (pq *priorityQueue[T]) insert(item T) {
	// TODO: increase by square when capacity is full
	pq.heap = append(pq.heap, item)
	pq.n++
	pq.swim(pq.n)
}

func (pq *priorityQueue[T]) delete() T {
	top := pq.top()
	pq.swap(1, pq.n)
	pq.heap = pq.heap[:pq.n]
	pq.n--
	pq.sink(1)
	return top
}

func (pq *priorityQueue[_]) size() int {
	return pq.n
}

func (pq *priorityQueue[_]) isEmpty() bool {
	return pq.n == 0
}

func (pq *priorityQueue[_]) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *priorityQueue[T]) swim(start int) {
	for k := start; k > 1 && pq.less(pq.heap[k/2], pq.heap[k]); k = k / 2 {
		pq.swap(k/2, k)
	}
}

func (pq *priorityQueue[T]) sink(k int) {
	// While k is a parent with some children
	for 2*k <= pq.n {

		// First child of a parent k
		j := 2 * k

		// If first child is less than second
		// we choose second one for exchange.
		// Parent should be swapped with biggest child
		if j < pq.n && pq.less(pq.heap[j], pq.heap[j+1]) {
			j++
		}

		// If parent is already bigger than biggest child
		// we found the position
		if !pq.less(pq.heap[k], pq.heap[j]) {
			break
		}

		// swap parent node with child
		pq.swap(j, k)

		// child node is a new parent
		k = j
	}
}
