package mathevaluator

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	result := Evaluate("( 1 + 1 )")

	if result != 2 {
		log.Fatalf("wrong answer: %v", result)
	}
}

func TestSqrt(t *testing.T) {
	result := Evaluate("sqrt ( 9 )")

	if result != 3 {
		log.Fatalf("wrong answer: %v", result)
	}
}

func TestComplex(t *testing.T) {
	result := Evaluate("( ( 1 + ( 2 - ( 3 * ( 4 / ( 1 + 1 ) ) ) ) ) + sqrt ( 9 ) )")

	if result != 0 {
		log.Fatalf("wrong answer: %v", result)
	}
}
