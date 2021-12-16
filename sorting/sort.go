package sorting

type SliceSorter[T any] func([]T, func(T, T) bool)
