package sorting

// TODO: compare function should receive pointers to slice elements
// to prevent unnecessary coping
type SliceSorter[T any] func([]T, func(T, T) bool)
