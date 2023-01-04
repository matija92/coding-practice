package graphdfs

import (
	"coding-practice/data-structures/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFSWithEmptyGraph(t *testing.T) {
	g := graph.NewUndirected()

	iterator, err := NewIterator(g)
	assert.Nil(t, err)

	// Creat fake start node
	nodeA := graph.NewNode("A", 0)
	res, err := iterator.Iterate(nodeA)
	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestDFSWithRootNode(t *testing.T) {

	// Input:
	// 		A

	// Output: A
	g := graph.NewUndirected()

	// Add nodes
	nodeA := graph.NewNode("A", 0)
	err := g.AddNode(nodeA)
	assert.Nil(t, err)

	iterator, err := NewIterator(g)
	assert.Nil(t, err)

	res, err := iterator.Iterate(nodeA)
	assert.Nil(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, []*graph.Node{nodeA}, res)
}

func TestSimpleDFS(t *testing.T) {

	// Input:
	// 		A
	// 	  /  \
	// 	 B    C
	// 	     / \
	//      D   F
	//     /
	//    E

	// Output: A -> B -> C -> D -> E -> F
	g := graph.NewUndirected()

	// Add nodes
	nodeA := graph.NewNode("A", 0)
	nodeB := graph.NewNode("B", 0)
	nodeC := graph.NewNode("C", 0)
	nodeD := graph.NewNode("D", 0)
	nodeE := graph.NewNode("E", 0)
	nodeF := graph.NewNode("F", 0)
	err := g.AddNode(nodeA)
	assert.Nil(t, err)
	err = g.AddNode(nodeB)
	assert.Nil(t, err)
	err = g.AddNode(nodeC)
	assert.Nil(t, err)
	err = g.AddNode(nodeD)
	assert.Nil(t, err)
	err = g.AddNode(nodeE)
	assert.Nil(t, err)
	err = g.AddNode(nodeF)
	assert.Nil(t, err)

	// Add edges
	err = g.AddEdge(nodeA, nodeB, 0)
	assert.Nil(t, err)
	err = g.AddEdge(nodeA, nodeC, 0)
	assert.Nil(t, err)

	err = g.AddEdge(nodeC, nodeD, 0)
	assert.Nil(t, err)
	err = g.AddEdge(nodeC, nodeF, 0)
	assert.Nil(t, err)

	err = g.AddEdge(nodeD, nodeE, 0)
	assert.Nil(t, err)

	iterator, err := NewIterator(g)
	assert.Nil(t, err)

	res, err := iterator.Iterate(nodeA)
	assert.Nil(t, err)
	assert.Len(t, res, 6)
	assert.Equal(t, []*graph.Node{nodeA, nodeB, nodeC, nodeD, nodeE, nodeF}, res)
}
