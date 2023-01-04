package djikstra

import (
	"coding-practice/pkg/data-structures/graph"
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

// Input graph
//
// [[0,1,4],[1,2,1],[2,3,8],[3,4,9],[3,5,5],[5,4,12],[6,5,7],[0,6,10],[1,6,2],[6,2,6]]

// Expected results
//
// Vertex Distance from Source:
// 0	0
// 1	4
// 2	5
// 3	13
// 4	22
// 5	13
// 6	6
func TestSimpleGraph(t *testing.T) {
	g := graph.NewDirected()

	nodes := createNodes(7)
	for _, node := range nodes {
		g.AddNode(node)
	}

	g.AddEdge(nodes[0], nodes[1], 4)
	g.AddEdge(nodes[0], nodes[6], 10)
	g.AddEdge(nodes[1], nodes[2], 1)
	g.AddEdge(nodes[1], nodes[6], 2)
	g.AddEdge(nodes[2], nodes[3], 8)
	g.AddEdge(nodes[3], nodes[4], 9)
	g.AddEdge(nodes[3], nodes[5], 5)
	g.AddEdge(nodes[5], nodes[4], 12)
	g.AddEdge(nodes[6], nodes[2], 6)
	g.AddEdge(nodes[6], nodes[5], 7)

	startNode := nodes[0]
	err := Run(g, startNode)
	assert.Nil(t, err)

	assert.Equal(t,
		[]int{
			0,
			4,
			5,
			13,
			22,
			13,
			6,
		},
		[]int{
			nodes[0].Value(),
			nodes[1].Value(),
			nodes[2].Value(),
			nodes[3].Value(),
			nodes[4].Value(),
			nodes[5].Value(),
			nodes[6].Value(),
		})
}
