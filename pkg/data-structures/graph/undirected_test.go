package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingEdges(t *testing.T) {
	graph := NewUndirected()

	nodeA := NewNode("A", 0)
	nodeB := NewNode("B", 0)

	// Adding edge before the nodes are added results in an error
	err := graph.AddEdge(nodeA, nodeB, 10)
	assert.NotNil(t, err)

	err = graph.AddNode(nodeA)
	assert.Nil(t, err)

	err = graph.AddEdge(nodeA, nodeB, 10)
	assert.NotNil(t, err)

	err = graph.AddNode(nodeB)
	assert.Nil(t, err)

	err = graph.AddEdge(nodeA, nodeB, 10)
	assert.Nil(t, err)

	// Adding existing edge results in an error
	err = graph.AddEdge(nodeA, nodeB, 10)
	assert.NotNil(t, err)
}

func TestSimpleGraph(t *testing.T) {
	graph := NewUndirected()

	nodeA := NewNode("A", 0)
	nodeB := NewNode("B", 0)

	err := graph.AddNode(nodeA)
	assert.Nil(t, err)

	err = graph.AddNode(nodeB)
	assert.Nil(t, err)

	err = graph.AddEdge(nodeA, nodeB, 10)
	assert.Nil(t, err)

	nodeAEdges := graph.Edges(nodeA)
	assert.Len(t, nodeAEdges, 1)

	nodeAEdge := nodeAEdges[0]
	assert.Equal(t, nodeA, nodeAEdge.N1)
	assert.Equal(t, nodeB, nodeAEdge.N2)
	assert.Equal(t, 10, nodeAEdge.Weight)

	nodeBEdges := graph.Edges(nodeB)
	assert.Len(t, nodeBEdges, 1)

	nodeBEdge := nodeBEdges[0]
	assert.Equal(t, nodeB, nodeBEdge.N1)
	assert.Equal(t, nodeA, nodeBEdge.N2)
	assert.Equal(t, 10, nodeAEdge.Weight)
}
