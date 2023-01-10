package mst

import (
	"coding-practice/data-structures/graph"
	"coding-practice/data-structures/heap"
)

func MinimumSpanningTree(g *graph.UndirectedGraph) []*graph.Edge {
	edges := make([]*graph.Edge, 0)
	added := make(map[*graph.Edge]bool)
	h := heap.New(heap.MinHeap)

	for _, node := range g.Nodes() {
		for _, edge := range g.Edges(node) {
			if !added[edge] {
				h.Push(edge)
				added[edge] = true
			}
		}
	}

	ufGraph := graph.NewUnionFind()
	for element, err := h.Pop(); err == nil && element != nil; element, err = h.Pop() {
		edge := element.(*graph.Edge)
		if !ufGraph.NodeExists(edge.N1) {
			ufGraph.AddNode(edge.N1)
		}

		if !ufGraph.NodeExists(edge.N2) {
			ufGraph.AddNode(edge.N2)
		}

		if !ufGraph.Connected(edge.N1, edge.N2) {
			ufGraph.AddEdge(edge.N1, edge.N2, edge.Weight)
			edges = append(edges, edge)
		}
	}

	return edges
}
