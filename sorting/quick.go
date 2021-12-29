package sorting

import (
	"math/rand"
	"time"
)

func exch[T any](items []T, a int, b int) {
	items[a], items[b] = items[b], items[a]
}

func partition[T any](items []T, lo int, hi int, less func(a, b T) bool) int {
	i, j := lo, hi+1
	v := items[lo]

	for {
		for i++; less(items[i], v); i++ {
			if i == hi {
				break
			}
		}

		for j--; less(v, items[j]); j-- {
			if j == lo {
				break
			}
		}

		if i >= j {
			break
		}

		exch(items, i, j)
	}
	exch(items, j, lo)

	return j
}

func doQuickSort[T any](items []T, lo int, hi int, less func(a, b T) bool) {
	if hi-lo <= 15 {
		Insertion(items[lo:hi+1], less)
		return
	}

	if lo >= hi {
		return
	}
	mi := partition(items, lo, hi, less)
	doQuickSort(items, lo, mi-1, less)
	doQuickSort(items, mi+1, hi, less)

}

func Quick[T any](items []T, less func(a, b T) bool) {

	// We shuffle array to prevent worst case scenario when array already sorted.
	// Without shuffling we get O(n^2) time. Another variant how to prevent
	// such performance drop is to compute median element in partition function.
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })
	doQuickSort(items, 0, len(items)-1, less)
}
