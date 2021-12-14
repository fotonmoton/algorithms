package sorting

import (
	"testing"
)

func TestSelection(t *testing.T) {
	CheckSorter(NewSelection())
}

func BenchmarkSelection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, NewSelection())
	}
}

func TestInsertion(t *testing.T) {
	CheckSorter(NewInsertion())
}

func BenchmarkInsertion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BenchmarkSort(10000, NewInsertion())
	}
}
