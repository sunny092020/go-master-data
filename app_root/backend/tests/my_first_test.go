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

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
			mid := (low + high) / 2
			if arr[mid] < x {
					low = mid + 1
			} else if arr[mid] > x {
					high = mid - 1
			} else {
					return mid
			}
	}

	return -1
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

	// assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

	// Assert that two strings are equal
	assert.Equal(t, "Hello, world!", "Hello, world!", "The two strings should be equal")

	// Assert that a condition is true
	assert.True(t, 1+1 == 2, "1+1 should equal 2")

	// Assert that a value is nil
	var x *int
	assert.Nil(t, x, "x should be nil")

	// Assert that a slice contains a specific element
	slice := []int{1, 2, 3, 4}
	assert.Contains(t, slice, 3, "slice should contain 3")
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Test a value in the array
	index := binarySearch(arr, 5)
	assert.Equal(t, 4, index, "5 should be at index 4")
	
	// Test a value not in the array
	index = binarySearch(arr, 20)
	assert.Equal(t, -1, index, "20 should not be found in the array")
	
	// Test an empty array
	index = binarySearch([]int{}, 5)
	assert.Equal(t, -1, index, "5 should not be found in an empty array")
}
