package stack

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	stack := NewStack[int]()

	stack.Push(1)
	stack.Push(2)

	if stack.Pop() != 2 {
		log.Fatal("wrong stack value")
	}
}

func TestSize(t *testing.T) {
	stack := NewStack[int]()

	if !stack.IsEmpty() {
		log.Fatal("should be empty")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		log.Fatal("wrong size")
	}

}
