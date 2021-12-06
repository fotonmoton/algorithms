package sorting

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

type Sorter interface {
	Sort(Sortable)
	// TODO: add generic slice sort when type variables are landed
	// SortSlice[T any](T, func(i, j int) bool)
}
