package unionfind

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func testUnion(i implementation) {
	qf := i.create(4)

	qf.Union(0, 1)

	if !qf.Connected(0, 1) {
		log.Fatalf("%v: sites should be connected after union", i.name)
	}

	if qf.Find(0) != qf.Find(1) {
		log.Fatalf("%v after union sites should be in the same component", i.name)
	}

	qf.Union(2, 3)

	if !qf.Connected(2, 3) {
		log.Fatalf("%v: sites should be connected after union", i.name)
	}

	if qf.Find(2) != qf.Find(3) {
		log.Fatalf("%v after union sites should be in the same component", i.name)
	}

	qf.Union(1, 2)

	if qf.Count() != 1 {
		log.Fatalf("%v after union count should be decreased", i.name)
	}
}

func testFind(i implementation) {
	qf := i.create(2)

	if qf.Find(0) != 0 || qf.Find(1) != 1 {
		log.Fatalf("%v Before union all sites belongs to component with same number", i.name)
	}
}

func testFile(fileName string, components int, i implementation) {
	lines := readByLine(fileName)
	count, _ := strconv.Atoi(<-lines)
	qf := i.create(count)

	for line := range lines {
		first, second := pair(line)
		if qf.Connected(first, second) {
			continue
		}
		qf.Union(first, second)
	}

	if components != qf.Count() {
		log.Fatalf("%v: Expected components count: %v, got: %v", i.name, components, qf.Count())
	}
	fmt.Printf("%v: Components: %v\n", i.name, qf.Count())
}

func readByLine(fileName string) <-chan string {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	line := make(chan string)

	go func() {
		defer close(line)
		for scanner.Scan() {
			line <- scanner.Text()
		}
	}()

	return line
}

func pair(str string) (int, int) {
	numbers := strings.Split(str, " ")

	a, _ := strconv.Atoi(numbers[0])
	b, _ := strconv.Atoi(numbers[1])

	return a, b
}

func testCount(i implementation) {
	qf := i.create(10)

	if qf.Count() != 10 {
		log.Fatalf("%v: Before any union number of components should be equal to number of sites", i.name)
	}
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
