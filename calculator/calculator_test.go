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
