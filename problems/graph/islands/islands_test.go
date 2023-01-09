package islands

import (
	"coding-practice/data-structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test graph

// A - B - C
//
// 	   D
//
//		   - G
//	E - F
//		   - H
func TestIslands(t *testing.T) {
	g := graph.NewUnionFind()

	nA, nB, nC, nD, nE, nF, nG, nH := graph.NewNode("A", 0), graph.NewNode("B", 0), graph.NewNode("C", 0), graph.NewNode("D", 0), graph.NewNode("E", 0), graph.NewNode("F", 0), graph.NewNode("G", 0), graph.NewNode("H", 0)
	g.AddNode(nA)
	g.AddNode(nB)
	g.AddNode(nC)
	g.AddNode(nD)
	g.AddNode(nE)
	g.AddNode(nF)
	g.AddNode(nG)
	g.AddNode(nH)

	g.AddEdge(nA, nB, 0)
	g.AddEdge(nB, nC, 0)
	g.AddEdge(nE, nF, 0)
	g.AddEdge(nF, nG, 0)
	g.AddEdge(nF, nH, 0)

	assert.Equal(t, 3, NoOfIslands(g))
}
