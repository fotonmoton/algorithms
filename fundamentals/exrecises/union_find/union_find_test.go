package unionfind

import (
	"testing"
)

type implementation struct {
	name   string
	create func(int) UnionFind
}

var implementations = []implementation{
	{
		name:   "quick find",
		create: NewQuickFind,
	},
	{
		name:   "quick union",
		create: NewQuickUnion,
	},
	{
		name:   "weighted quick union",
		create: NewWeightedQuickUnion,
	},
}

func TestCount(t *testing.T) {
	for _, implementation := range implementations {
		testCount(implementation)
	}
}

func TestFind(t *testing.T) {
	for _, implementation := range implementations {
		testFind(implementation)
	}
}

func TestUnion(t *testing.T) {
	for _, implementation := range implementations {
		testUnion(implementation)
	}
}

func BenchmarkLarge(b *testing.B) {
	testFile("largeUF.txt", 6, implementation{
		name:   "weighted quick union",
		create: NewWeightedQuickUnion,
	})
}

func BenchmarkTiny(b *testing.B) {
	for _, implementation := range implementations {
		testFile("tinyUF.txt", 2, implementation)
	}
}

func BenchmarkMedium(b *testing.B) {
	for _, implementation := range implementations {
		testFile("mediumUF.txt", 3, implementation)
	}
}
