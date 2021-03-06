package queue

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	queue := NewQueue[int]()

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	first, second := queue.Dequeue(), queue.Dequeue()

	if first != 10 && second != 20 {
		log.Fatal("wrong order")
	}
}

func TestSize(t *testing.T) {
	queue := NewQueue[int]()

	if queue.Size() != 0 {
		log.Fatal("empty queue should have size 0")
	}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	if queue.Size() != 3 {
		log.Fatal("wrong size")
	}
}
