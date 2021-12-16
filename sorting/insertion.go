package sorting

type insertion struct{}

func Insertion[T any](items []T, less func(a, b T) bool) {
	len := len(items)
	for i := 1; i < len; i++ {
		for j := i; j > 0 && less(items[j], items[j-1]); j-- {
			items[j], items[j-1] = items[j-1], items[j]
		}
	}
}
