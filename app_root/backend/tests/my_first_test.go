package my_first_test

import (
	"container/list"
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

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	data     int
	visited  bool
	adjacent []*Vertex
}

func NewGraph() *Graph {
	return &Graph{vertices: []*Vertex{}}
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) BFS(start *Vertex) []*Vertex {
	queue := list.New()
	var visitedVertices []*Vertex

	queue.PushBack(start)
	start.visited = true

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		vertex := element.Value.(*Vertex)
		visitedVertices = append(visitedVertices, vertex)

		for _, v := range vertex.adjacent {
			if !v.visited {
				v.visited = true
				queue.PushBack(v)
			}
		}
	}

	return visitedVertices
}

func TestBFS(t *testing.T) {
	graph := NewGraph()

	// create vertices
	vertex1 := &Vertex{data: 1}
	vertex2 := &Vertex{data: 2}
	vertex3 := &Vertex{data: 3}
	vertex4 := &Vertex{data: 4}
	vertex5 := &Vertex{data: 5}

	// add edges
	vertex1.adjacent = []*Vertex{vertex2, vertex3}
	vertex2.adjacent = []*Vertex{vertex4, vertex5}
	vertex3.adjacent = []*Vertex{vertex5}

	// add vertices to graph
	graph.AddVertex(vertex1)
	graph.AddVertex(vertex2)
	graph.AddVertex(vertex3)
	graph.AddVertex(vertex4)
	graph.AddVertex(vertex5)

	// perform BFS
	visitedVertices := graph.BFS(vertex1)

	// check that all vertices were visited
	assert.Equal(t, len(visitedVertices), 5)

	// check the order in which vertices were visited
	assert.Equal(t, visitedVertices[0].data, 1)
	assert.Equal(t, visitedVertices[1].data, 2)
	assert.Equal(t, visitedVertices[2].data, 3)
	assert.Equal(t, visitedVertices[3].data, 4)
	assert.Equal(t, visitedVertices[4].data, 5)
}
