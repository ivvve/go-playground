package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var addTests = []struct {
	nums   []int
	result int
}{
	{[]int{1}, 1},
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3}, 6},
	{[]int{1, 2, 3, 4}, 10},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		assert.Equal(t, Add(test.nums...), test.result)
	}
}

var subtractTests = []struct {
	nums   []int
	result int
}{
	{[]int{1, 2}, -1},
	{[]int{2, 1}, 1},
}

func TestSubtract(t *testing.T) {
	for _, test := range subtractTests {
		assert.Equal(t, Subtract(test.nums[0], test.nums[1]), test.result)
	}
}

var multiplyTests = []struct {
	nums   []int
	result int
}{
	{[]int{2, 1}, 2},
	{[]int{2, -2}, -4},
}

func TestMultiply(t *testing.T) {
	for _, test := range multiplyTests {
		assert.Equal(t, Multiply(test.nums[0], test.nums[1]), test.result)
	}
}

var divideTests = []struct {
	nums     []int
	result   float64
	hasError bool
}{
	{[]int{2, 1}, 2.0, false},
	{[]int{1, 2}, 0.5, false},
	{[]int{1, 0}, 0, true},
}

func TestDivide(t *testing.T) {
	for _, test := range divideTests {
		result, err := Divide(test.nums[0], test.nums[1])

		if test.hasError {
			assert.NotNil(t, err)
		} else {
			assert.Equal(t, result, test.result)
		}
	}
}
