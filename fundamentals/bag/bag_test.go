package bag

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	bag := NewBag[int]()
	sum := 0

	bag.Add(1)
	bag.Add(2)
	bag.Add(3)

	bag.ForEach(func(i int) { sum += i })

	if sum != 6 {
		log.Fatal("wrong items in bag")
	}
}

func TestEmpty(t *testing.T) {
	bag := NewBag[int]()
	sum := 0

	if bag.Size() != 0 || !bag.IsEmpty() {
		log.Fatal("bag should be empty")
	}

	bag.ForEach(func(i int) { sum += i })

	if sum != 0 {
		log.Fatal("wrong items in bag")
	}
}
