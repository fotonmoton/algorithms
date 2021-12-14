package sorting

type shell struct{}

func (*shell) Sort(items Sortable) {
	len := items.Len()
	gap := 1

	// Calculating gap maximum value.
	// This is for "Sedgewick gap sequence" variation.
	// Another sequences can be used
	for gap < len/3 {
		gap = gap*3 + 1
	}

	// This loop needed to progressively degrease gap until siple insertion
	// sort will be used
	for gap >= 1 {

		// Insertion sort loop
		for i := gap; i < len; i++ {

			// Instead of comparing adjacent elements we compare
			// gap distance elements and swap them
			for j := i; j >= gap && items.Less(j, j-gap); j -= gap {
				items.Swap(j, j-gap)
			}
		}

		// "Sedgewick gap sequence"
		gap = gap / 3
	}
}

func NewShell() Sorter {
	return &shell{}
}
