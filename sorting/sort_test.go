package sorting

import (
	"testing"
)

func TestSelectionSlice(t *testing.T) {
	CheckSliceSorter(Selection[int])
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, Selection[int])
	}
}

func TestInsertion(t *testing.T) {
	CheckSliceSorter(Insertion[int])
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, Insertion[int])
	}
}

func TestShell(t *testing.T) {
	CheckSliceSorter(Shell[int])
}

func BenchmarkShell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, Shell[int])
	}
}

func TestMerge(t *testing.T) {
	CheckSliceSorter(Merge[int])
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, Merge[int])
	}
}

func TestParallelMerge(t *testing.T) {
	CheckSliceSorter(ParallelMerge[int])
}

func BenchmarkParallelMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, ParallelMerge[int])
	}
}

func TestBottomUpMerge(t *testing.T) {
	CheckSliceSorter(BottomUpMerge[int])
}

func BenchmarkBottomUpMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, BottomUpMerge[int])
	}
}

func TestQuick(t *testing.T) {
	CheckSliceSorter(Quick[int])
}

func BenchmarkQuick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, Quick[int])
	}
}
