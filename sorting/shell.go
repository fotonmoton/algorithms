package sorting

func Shell[T any](items []T, less func(a, b T) bool) {
	len := len(items)
	gap := 1

	// Calculating gap maximum value.
	// This is for "Sedgewick gap sequence" variation.
	// Another sequences can be used
	for gap < len/3 {
		gap = gap*3 + 1
	}

	// This loop needed to progressively decrease gap until simple insertion
	// sort will be used
	for gap >= 1 {

		// Insertion sort loop
		for i := gap; i < len; i++ {

			// Instead of comparing adjacent elements we compare
			// gap distance elements and swap them
			for j := i; j >= gap && less(items[j], items[j-gap]); j -= gap {
				items[j], items[j-gap] = items[j-gap], items[j]
			}
		}

		// "Sedgewick gap sequence"
		gap = gap / 3
	}
}
