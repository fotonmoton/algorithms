package sorting

func Selection[T any](items []T, less func(a, b T) bool) {
	len := len(items)
	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if less(items[j], items[min]) {
				min = j
			}
		}
		items[min], items[i] = items[i], items[min]
	}
}
