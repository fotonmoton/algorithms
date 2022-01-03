package sorting

func swap[T any](i, j int, items []T) {
	items[i], items[j] = items[j], items[i]
}

func swim[T any](child int, depth int, items []T, less func(T, T) bool) {
	for {
		parent := (child - 1) / 2

		if child <= 0 || less(items[child], items[parent]) {
			break
		}

		swap(parent, child, items)

		child = parent
	}
}

func sink[T any](parent int, depth int, items []T, less func(T, T) bool) {
	for {
		child := parent*2 + 1

		if child >= depth {
			break
		}

		if child+1 < depth && less(items[child], items[child+1]) {
			child++
		}

		if !less(items[parent], items[child]) {
			break
		}

		swap(parent, child, items)

		parent = child
	}
}

func Heap[T any](items []T, less func(T, T) bool) {
	len := len(items)

	for k := len / 2; k >= 0; k-- {
		sink(k, len, items, less)
	}

	for len > 0 {
		len--
		swap(0, len, items)
		sink(0, len, items, less)
	}
}
