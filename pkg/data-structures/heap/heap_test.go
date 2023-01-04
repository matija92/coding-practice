package heap

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function for generating and array of n random integers (0-9)
func gen(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(10)
	}
	return arr
}

func TestEmptyHeap(t *testing.T) {
	heap := Heap{}

	// Test Peek()
	val, err := heap.Peek()
	assert.Equal(t, 0, val)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrHeapEmpty))

	// Test Pop()
	val, err = heap.Pop()
	assert.Equal(t, 0, val)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrHeapEmpty))
}

func TestDefaultHeap(t *testing.T) {
	arr := []int{5, 12, -1, 4, 3, 9, 20}
	sortedArr := []int{-1, 3, 4, 5, 9, 12, 20}

	// Initialize min heap (implicit)
	heap := Heap{}

	// Populate unsorted array
	for _, el := range arr {
		heap.Push(el)
	}

	res := []int{}
	for el, err := heap.Pop(); err == nil; el, err = heap.Pop() {
		res = append(res, el)
	}

	assert.Len(t, res, 7)
	assert.Equal(t, sortedArr, res)
}

func TestHeapModes(t *testing.T) {
	arr := []int{5, 12, -1, 4, 3, 9, 20}
	sortedAsc := []int{-1, 3, 4, 5, 9, 12, 20}
	sortedDesc := []int{20, 12, 9, 5, 4, 3, -1}

	// Initialize min heap
	heap := New(MinHeap)

	// Populate unsorted array
	for _, el := range arr {
		heap.Push(el)
	}

	res := []int{}
	for el, err := heap.Pop(); err == nil; el, err = heap.Pop() {
		res = append(res, el)
	}

	assert.Len(t, res, 7)
	assert.Equal(t, sortedAsc, res)

	// Initialize max heap
	heap = New(MaxHeap)

	// Populate unsorted array
	for _, el := range arr {
		heap.Push(el)
	}

	res = []int{}
	for el, err := heap.Pop(); err == nil; el, err = heap.Pop() {
		res = append(res, el)
	}

	assert.Len(t, res, 7)
	assert.Equal(t, sortedDesc, res)
}

func TestHeapFromArray(t *testing.T) {
	arr := []int{5, 12, -1, 4, 3, 9, 20}
	sortedAsc := []int{-1, 3, 4, 5, 9, 12, 20}
	sortedDesc := []int{20, 12, 9, 5, 4, 3, -1}

	// Initialize min heap
	heap := NewFromArray(arr, MinHeap)

	res := make([]int, 0)
	for el, err := heap.Pop(); err == nil; el, err = heap.Pop() {
		res = append(res, el)
	}

	assert.Len(t, res, 7)
	assert.Equal(t, sortedAsc, res)

	// Reset array
	arr = []int{5, 12, -1, 4, 3, 9, 20}

	// Initialize max heap
	heap = NewFromArray(arr, MaxHeap)

	res = make([]int, 0)
	for el, err := heap.Pop(); err == nil; el, err = heap.Pop() {
		res = append(res, el)
	}

	assert.Len(t, res, 7)
	assert.Equal(t, sortedDesc, res)
}
