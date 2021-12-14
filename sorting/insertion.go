package sorting

type insertion struct{}

func (*insertion) Sort(items Sortable) {
	len := items.Len()
	for i := 1; i < len; i++ {
		for j := i; j > 0 && items.Less(j, j-1); j-- {
			items.Swap(j, j-1)
		}
	}
}

func NewInsertion() Sorter {
	return &insertion{}
}
