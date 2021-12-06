package sorting

type selection struct{}

func (*selection) Sort(items Sortable) {
	len := items.Len()
	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if items.Less(j, min) {
				min = j
			}
		}
		items.Swap(min, i)
	}
}

func NewSelection() Sorter {
	return &selection{}
}
