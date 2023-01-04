package heap

import (
	"fmt"
	"math"
	"strings"
)

type heapError string

func (e heapError) Error() string {
	return string(e)
}

const (
	ErrHeapEmpty heapError = "heap is empty"
)

type heapMode int

const (
	MinHeap heapMode = iota
	MaxHeap
)

// Heap can be initialized either by implicit allocaction: heap := &Heap{}
// Or using the New(mode heapMode) constructor
// By default Heap is in MinHeap mode
type Heap struct {
	mode heapMode
	arr  []int
}

// Constructor for empty Heap, providing Heap mode
func New(mode heapMode) *Heap {
	return &Heap{
		mode: mode,
		arr:  make([]int, 0),
	}
}

// Constructor for new Heap, using unsorted array and providing Heap mode
// The array provided is owned by the underlying heap object after calling this constructor
func NewFromArray(arr []int, mode heapMode) *Heap {
	heap := &Heap{
		mode: mode,
		arr:  arr,
	}

	n := len(heap.arr)
	if n > 0 {
		lastNode := n - 1
		lastNonLeafNode := (lastNode - 1) / 2

		for i := lastNonLeafNode; i >= 0; i-- {
			heap.heapifyTD(i)
		}

	}

	return heap
}

// Push element on to the heap
func (h *Heap) Push(a int) {
	h.arr = append(h.arr, a)
	h.heapifyBU(len(h.arr) - 1)
}

// Remove the root element from the heap, and rebalance
func (h *Heap) Pop() (int, error) {
	if len(h.arr) == 0 {
		return 0, ErrHeapEmpty
	}

	res := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.heapifyTD(0)

	return res, nil
}

// Peek looks at the root value of the heap, but doesn't remove it
func (h *Heap) Peek() (int, error) {
	if len(h.arr) == 0 {
		return 0, ErrHeapEmpty
	}

	return h.arr[0], nil
}

// Helper function for printing out heap contents
func (h *Heap) Dump() {
	fmt.Println("")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("Heap Stats:")

	rows := int(math.Round(math.Log2(float64(len(h.arr)))))
	index := 0
	for i := 0; i < rows; i++ {
		noEl := int(math.Pow(2.0, float64(i)))
		endIndex := min(index+noEl, len(h.arr)-1)

		preTab := (rows - i - 1) * 2
		tmpArr := []string{}
		for _, el := range h.arr[index:endIndex] {
			tmpArr = append(tmpArr, fmt.Sprintf("%d", el))
		}
		fmt.Println(indent(preTab, strings.Join(tmpArr, indent(rows-i, ""))))
		index = endIndex
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-=")
	fmt.Println("")
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) heapifyTD(i int) {
	nodeIndex := i
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2

	if nodeIndex < 0 || nodeIndex >= len(h.arr) {
		return
	}

	target := nodeIndex
	switch h.mode {
	case MaxHeap:
		if leftIndex < len(h.arr) && h.arr[leftIndex] > h.arr[target] {
			target = leftIndex
		}
		if rightIndex < len(h.arr) && h.arr[rightIndex] > h.arr[target] {
			target = rightIndex
		}
	case MinHeap:
		if leftIndex < len(h.arr) && h.arr[leftIndex] < h.arr[target] {
			target = leftIndex
		}
		if rightIndex < len(h.arr) && h.arr[rightIndex] < h.arr[target] {
			target = rightIndex
		}
	}

	if target != nodeIndex {
		h.swap(nodeIndex, target)
		h.heapifyTD(target)
	}

}

func (h *Heap) heapifyBU(i int) {
	nodeIndex := i
	parentIndex := (i - 1) / 2

	if nodeIndex == parentIndex || i == 0 {
		return
	}

	switch h.mode {
	case MaxHeap:
		if nodeIndex != parentIndex && h.arr[parentIndex] < h.arr[nodeIndex] {
			h.swap(parentIndex, nodeIndex)
			h.heapifyBU(parentIndex)
		}
	case MinHeap:
		if nodeIndex != parentIndex && h.arr[parentIndex] > h.arr[nodeIndex] {
			h.swap(parentIndex, nodeIndex)
			h.heapifyBU(parentIndex)
		}
	}

}

// Unoptimized version of top down heapify
func (h *Heap) heapifyTD_Old(i int) {
	nodeIndex := i
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2

	if nodeIndex < 0 || nodeIndex >= len(h.arr) {
		return
	}

	switch h.mode {
	case MaxHeap:
		if leftIndex < len(h.arr) && h.arr[leftIndex] > h.arr[nodeIndex] {
			h.swap(leftIndex, nodeIndex)
			h.heapifyTD_Old(leftIndex)
		}
		if rightIndex < len(h.arr) && h.arr[rightIndex] > h.arr[nodeIndex] {
			h.swap(rightIndex, nodeIndex)
			h.heapifyTD_Old(rightIndex)
		}
	case MinHeap:
		if leftIndex < len(h.arr) && h.arr[leftIndex] < h.arr[nodeIndex] {
			h.swap(leftIndex, nodeIndex)
			h.heapifyTD_Old(leftIndex)
		}
		if rightIndex < len(h.arr) && h.arr[rightIndex] < h.arr[nodeIndex] {
			h.swap(rightIndex, nodeIndex)
			h.heapifyTD_Old(rightIndex)
		}
	}

}

func indent(n int, s string) string {
	res := ""
	for i := 0; i < n; i++ {
		res += " "
	}
	res += s
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
