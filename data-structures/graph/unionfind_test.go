package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Nodes: A B C D E
// Add edge A -> B
// Add edge C -> D
// Add edge B -> D
// Add edge A -> E
func TestUnionFind(t *testing.T) {
	graph := NewUnionFind()

	nodeA := NewNode("A", 0)
	nodeB := NewNode("B", 0)
	nodeC := NewNode("C", 0)
	nodeD := NewNode("D", 0)
	nodeE := NewNode("E", 0)

	err := graph.AddNode(nodeA)
	assert.Nil(t, err)
	err = graph.AddNode(nodeB)
	assert.Nil(t, err)
	err = graph.AddNode(nodeC)
	assert.Nil(t, err)
	err = graph.AddNode(nodeD)
	assert.Nil(t, err)
	err = graph.AddNode(nodeE)
	assert.Nil(t, err)

	// Validate root nodes are equal to nodes
	rootA := graph.Find(nodeA)
	assert.Equal(t, nodeA, rootA)

	rootB := graph.Find(nodeB)
	assert.Equal(t, nodeB, rootB)

	rootC := graph.Find(nodeC)
	assert.Equal(t, nodeC, rootC)

	rootD := graph.Find(nodeD)
	assert.Equal(t, nodeD, rootD)

	rootE := graph.Find(nodeE)
	assert.Equal(t, nodeE, rootE)

	// Adding A -> B
	err = graph.AddEdge(nodeA, nodeB, 0)
	assert.Nil(t, err)

	// Adding C -> D
	err = graph.AddEdge(nodeC, nodeD, 0)
	assert.Nil(t, err)

	// Adding B -> D
	err = graph.AddEdge(nodeB, nodeD, 0)
	assert.Nil(t, err)

	// Adding A -> E
	err = graph.AddEdge(nodeA, nodeE, 0)
	assert.Nil(t, err)

	// Validate root nodes are the same
	rootNode := graph.Find(nodeA)
	rootB = graph.Find(nodeB)
	rootC = graph.Find(nodeC)
	rootD = graph.Find(nodeD)
	rootE = graph.Find(nodeE)

	assert.Equal(t, rootNode, rootB)
	assert.Equal(t, rootNode, rootC)
	assert.Equal(t, rootNode, rootD)
	assert.Equal(t, rootNode, rootE)
}
