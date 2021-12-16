package sorting

func merge[T any](items []T, low, mid, hig int, less func(a, b T) bool, aux []T) {
	i, j := low, mid+1

	for k := low; k <= hig; k++ {
		aux[k] = items[k]
	}

	for k := low; k <= hig; k++ {
		if i > mid {
			items[k] = aux[j]
			j++
		} else if j > hig {
			items[k] = aux[i]
			i++
		} else if less(aux[i], aux[j]) {
			items[k] = aux[i]
			i++
		} else {
			items[k] = aux[j]
			j++
		}
	}

}

func doSort[T any](items []T, low, hig int, less func(a, b T) bool, aux []T) {

	// Optimization 2: if array length is less then 15
	// it's imperically prooven that insertion sort
	// for such arrays can speed up merge sort up to 15%
	if hig-low <= 15 {
		Insertion(items[low:hig+1], less)
		return
	}

	if hig <= low {
		return
	}

	mid := low + (hig-low)/2

	doSort(items, low, mid, less, aux)
	doSort(items, mid+1, hig, less, aux)

	// Optimization 1: if two subarrays already sorted
	// we can skip merging them
	if less(items[mid], items[mid+1]) {
		return
	}

	merge(items, low, mid, hig, less, aux)
}

func Merge[T any](items []T, less func(a, b T) bool) {

	len := len(items)

	doSort(items, 0, len-1, less, make([]T, len))
}
