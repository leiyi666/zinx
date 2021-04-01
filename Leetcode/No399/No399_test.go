package No399

import (
	"fmt"
	"testing"
)

func TestCalcEquation(t *testing.T) {
	/*equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}*/
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
		{"bc", "cd"},
	}
	values := []float64{1.5, 2.5, 5.0}
	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}
	result := calcEquation(equations, values, queries)
	fmt.Println("result: ", result)
}
