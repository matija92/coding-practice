package djikstra

import (
	"coding-practice/pkg/data-structures/graph"
	"coding-practice/pkg/data-structures/heap"
	"errors"
)

// Run assumes all nodes and edges have been added to the graph
// The functions assumes all node values are set to MaxInt32
func Run(g graph.Graph, n *graph.Node) error {
	if g == nil || n == nil {
		return errors.New("invalid input")
	}

	if !g.NodeExists(n) {
		return errors.New("node doesn't exist")
	}
	n.SetValue(0)

	found := make(map[*graph.Node]bool)
	done := make(map[*graph.Node]bool)
	heap := heap.New(heap.MinHeap)
	heap.Push(n)

	for element, err := heap.Pop(); err == nil; element, err = heap.Pop() {
		node := element.(*graph.Node)

		for _, edge := range g.Edges(node) {

			node2 := edge.N2

			if !found[node2] {
				heap.Push(node2)
				found[node2] = true
			}

			if node.Value()+edge.Weight < node2.Value() {
				node2.SetValue(node.Value() + edge.Weight)
			}

		}

		done[node] = true

	}

	return nil
}
