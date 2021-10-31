package fib

import (
	"testing"
)

var inputParams = []int{0, 1, 2, 3, 4, 5, 6, 7}
var expectedRes = []int{0, 1, 1, 2, 3, 5, 8, 13}

var functionsToTest = []func(int) int{Rec, Memo, Iterative}

func TestAll(t *testing.T) {
	for _, calculateFib := range functionsToTest {
		for i := 0; i < len(inputParams); i++ {
			actual := calculateFib(inputParams[i])
			expected := expectedRes[i]
			if actual != expected {
				t.Errorf("Result was incorrect, got: %d, expected: %d.", actual, expected)
			}
		}
	}
}
