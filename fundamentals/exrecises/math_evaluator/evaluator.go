package mathevaluator

import (
	"math"
	"strconv"
	"strings"

	"github.com/fotonmoton/algorithms/fundamentals/stack"
)

// Dijkstra's two stack algorithm for equation evaluation
func Evaluate(exression string) int {
	operations := stack.NewStack[string]()
	values := stack.NewStack[int]()

	for _, part := range strings.Split(exression, " ") {
		isOperation := part == "+" || part == "-" || part == "*" || part == "/" || part == "sqrt"

		if part == "(" {
			continue
		} else if part == ")" {
			lastOperation := operations.Pop()
			lastValue := values.Pop()

			switch lastOperation {
			case "+":
				lastValue = values.Pop() + lastValue
			case "-":
				lastValue = values.Pop() - lastValue
			case "*":
				lastValue = values.Pop() * lastValue
			case "/":
				lastValue = values.Pop() / lastValue
			case "sqrt":
				lastValue = int(math.Sqrt(float64(lastValue)))
			}
			values.Push(lastValue)
		} else if isOperation {
			operations.Push(part)
		} else {
			newValue, _ := strconv.Atoi(part)
			values.Push(newValue)
		}
	}

	return values.Pop()
}
