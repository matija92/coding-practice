package djikstra

import (
	"coding-practice/pkg/data-structures/graph"
	"math"
	"testing"
)

// Input graph
// 7 0
// [[0,1,4],[1,2,1],[2,3,8],[3,4,9],[3,5,5],[5,4,12],[6,5,7],[0,6,10],[1,6,2],[6,2,6]]

// Expected results
// Vertex Distance from Source:
// 0	0
// 1	4
// 2	5
// 3	13
// 4	22
// 5	13
// 6	6

func createGraph() (*graph.Node, graph.Graph) {
	g := graph.NewDirected()

	// Nodes
	n0 := graph.NewNode("0", math.MaxInt)
	n1 := graph.NewNode("1", math.MaxInt)
	n2 := graph.NewNode("2", math.MaxInt)
	n3 := graph.NewNode("3", math.MaxInt)
	n4 := graph.NewNode("4", math.MaxInt)
	n5 := graph.NewNode("5", math.MaxInt)
	n6 := graph.NewNode("6", math.MaxInt)
	g.AddNode(n0)
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)
	g.AddNode(n5)
	g.AddNode(n6)

	g.AddEdge(n0, n1, 4)
	g.AddEdge(n0, n6, 10)
	g.AddEdge(n1, n2, 1)
	g.AddEdge(n1, n6, 2)
	g.AddEdge(n2, n3, 8)
	g.AddEdge(n3, n4, 9)
	g.AddEdge(n3, n5, 5)
	g.AddEdge(n5, n4, 12)
	g.AddEdge(n6, n2, 6)
	g.AddEdge(n6, n5, 7)

	return n0, g
}

func TestSimpleGraph(t *testing.T) {
	startNode, g := createGraph()

	Run(g, startNode)
}
