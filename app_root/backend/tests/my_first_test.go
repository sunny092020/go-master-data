package my_first_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
			input    int
			expected int
	}{
			{2, 4},
			{-1, 1},
			{0, 2},
			{-5, -3},
			{99999, 100001},
	}

	for _, test := range tests {
			assert.Equal(Calculate(test.input), test.expected)
	}
}

func TestSomething(t *testing.T) {
  assert.True(t, true, "True is true!")
}

func TestSomethingMore(t *testing.T) {
	assert.Equal(t, 1, 1)
}
