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
	// Actually this is "less" function.
	// First element not necessary is less than second,
	// Can be used both ways to get reversed priority
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
	// After all operations "biggest" element
	// should be always at the beginning of the array
	return pq.heap[1]
}

func (pq *priorityQueue[T]) insert(item T) {
	// We can ignore "resize optimization":
	// append function will create new array for new slice
	// with doubled capacity when needed
	// https://github.com/golang/go/blob/master/src/runtime/slice.go#L188
	// https://go.dev/blog/slices-intro
	// https://go.dev/play/p/OKtCFskbp2t
	// https://stackoverflow.com/questions/23531737/how-the-slice-is-enlarged-by-append-is-the-capacity-always-doubled
	pq.heap = append(pq.heap, item)
	pq.n++
	pq.swim(pq.n)
}

func (pq *priorityQueue[T]) delete() T {
	top := pq.top()
	pq.swap(1, pq.n)
	// Discard "biggest" element from queue
	// for possible GC
	pq.heap = pq.heap[:pq.n]
	pq.n--
	// Reshape heap from top to bottom
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

func (pq *priorityQueue[T]) swim(child int) {
	// Until we reach top of the heap
	// and parent node is less than current child
	for child > 1 && pq.less(pq.heap[child/2], pq.heap[child]) {

		// We swap parent with the child
		pq.swap(child/2, child)

		// Parent node becomes new child
		// for next iteration
		child = child / 2
	}
}

func (pq *priorityQueue[T]) sink(parent int) {
	// While parent has some children
	for 2*parent <= pq.n {

		// First child of a parent
		child := 2 * parent

		// If first child is less than second
		// we choose second one for exchange.
		// Parent should be swapped with biggest child
		if child < pq.n && pq.less(pq.heap[child], pq.heap[child+1]) {
			child++
		}

		// If parent is already bigger than biggest child
		// we found the position
		if !pq.less(pq.heap[parent], pq.heap[child]) {
			break
		}

		// swap parent node with child
		pq.swap(parent, child)

		// child node is a new parent
		parent = child
	}
}
