package sorting

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

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

func CheckSorter(s Sorter) {
	rand.Seed(time.Now().Unix())

	actual := rand.Perm(1000)
	expected := make([]int, len(actual))
	copy(expected, actual)

	s.Sort(IntSort(actual))
	sort.Sort(IntSort(expected))

	if !SameInts(actual, expected) {
		log.Fatalf("wrong order:\n actual:\t%v\n expected:\t%v\n", actual, expected)
	}
}

func BenchmarkSort(numItems int, s Sorter) {
	rand.Seed(time.Now().Unix())
	items := rand.Perm(numItems)
	s.Sort(IntSort(items))
}
