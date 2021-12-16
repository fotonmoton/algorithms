package sorting

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

type IntSort []int

func (items IntSort) Len() int {
	return len(items)
}

func (items IntSort) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items IntSort) Less(i, j int) bool {
	return items[i] < items[j]
}

func SameInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func intCmp(a, b int) bool { return a < b }

func CheckSliceSorter(sorter SliceSorter[int]) {
	rand.Seed(time.Now().Unix())

	actual := rand.Perm(1000)
	expected := make([]int, len(actual))
	copy(expected, actual)

	sorter(actual, intCmp)
	sort.Sort(IntSort(expected))

	if !SameInts(actual, expected) {
		log.Fatalf("wrong order:\n actual:\t%v\n expected:\t%v\n", actual, expected)
	}
}

func BenchmarkSort(numItems int, sorter SliceSorter[int]) {
	rand.Seed(time.Now().Unix())
	items := rand.Perm(numItems)
	sorter(items, intCmp)
}
