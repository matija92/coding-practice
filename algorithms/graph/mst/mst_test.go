package mst

import (
	"coding-practice/data-structures/graph"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createNodes(n int) []*graph.Node {
	nodes := []*graph.Node{}
	for i := 0; i < n; i++ {
		nodes = append(nodes, graph.NewNode(fmt.Sprintf("%d", i), math.MaxInt))
	}
	return nodes
}

func TestMST(t *testing.T) {
	n := createNodes(9)

	g := graph.NewUndirected()

	for _, node := range n {
		g.AddNode(node)
	}

	g.AddEdge(n[0], n[1], 4)
	g.AddEdge(n[0], n[7], 8)
	g.AddEdge(n[1], n[7], 11)
	g.AddEdge(n[1], n[2], 8)
	g.AddEdge(n[2], n[8], 2)
	g.AddEdge(n[2], n[3], 7)
	g.AddEdge(n[2], n[5], 4)
	g.AddEdge(n[3], n[5], 14)
	g.AddEdge(n[3], n[4], 9)
	g.AddEdge(n[5], n[4], 10)
	g.AddEdge(n[5], n[6], 2)
	g.AddEdge(n[8], n[6], 6)
	g.AddEdge(n[7], n[6], 1)
	g.AddEdge(n[7], n[8], 7)

	edges := MinimumSpanningTree(g)
	assert.Len(t, edges, len(n)-1)

	sum := 0
	for _, e := range edges {
		sum += e.Weight
	}
	assert.Equal(t, 37, sum)
}
