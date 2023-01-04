package graphdfs

import (
	"coding-practice/pkg/data-structures/graph"
	"errors"
)

type dfsIterator struct {
	g       graph.Graph
	visited map[*graph.Node]bool
	result  []*graph.Node
}

func NewIterator(g graph.Graph) (*dfsIterator, error) {
	if g == nil {
		return nil, errors.New("empty input")
	}

	return &dfsIterator{
		g:       g,
		visited: make(map[*graph.Node]bool),
		result:  make([]*graph.Node, 0),
	}, nil
}

func (i *dfsIterator) Iterate(n *graph.Node) ([]*graph.Node, error) {
	if !i.g.NodeExists(n) {
		return nil, errors.New("node not a part of the graph")
	}

	if i.visited[n] {
		return nil, nil
	}

	i.result = append(i.result, n)
	i.visited[n] = true
	edges := i.g.Edges(n)
	for _, edge := range edges {
		if _, err := i.Iterate(edge.N2); err != nil {
			return nil, err
		}
	}

	return i.result, nil
}
