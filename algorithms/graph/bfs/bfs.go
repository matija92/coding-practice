package graphbfs

import (
	"coding-practice/data-structures/graph"
	"errors"
)

type bfsIterator struct {
	g       graph.Graph
	visited map[*graph.Node]bool
	result  []*graph.Node
}

func NewIterator(g graph.Graph) (*bfsIterator, error) {
	if g == nil {
		return nil, errors.New("empty input")
	}

	return &bfsIterator{
		g:       g,
		visited: make(map[*graph.Node]bool),
		result:  make([]*graph.Node, 0),
	}, nil
}

func (i *bfsIterator) Iterate(n *graph.Node) ([]*graph.Node, error) {
	if !i.g.NodeExists(n) {
		return nil, errors.New("node not a part of the graph")
	}

	queue := []*graph.Node{n}
	i.visited[n] = true

	var current *graph.Node
	for len(queue) > 0 {
		// Remove current element
		current = queue[0]
		queue = queue[1:]

		// Visit all unvisited neighbours
		edges := i.g.Edges(current)
		for _, edge := range edges {
			if !i.visited[edge.N2] {
				queue = append(queue, edge.N2)
			}
		}

		i.visited[current] = true
		i.result = append(i.result, current)
	}

	return i.result, nil
}
