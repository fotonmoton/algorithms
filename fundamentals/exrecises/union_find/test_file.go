package unionfind

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
