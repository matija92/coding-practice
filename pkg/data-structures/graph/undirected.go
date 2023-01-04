package graph

import "errors"

type UndirectedGraph struct {
	nodes map[*Node]bool // Lookup map for nodes
	edges map[*Node][]*edge
}

func NewUndirected() *UndirectedGraph {
	return &UndirectedGraph{
		nodes: make(map[*Node]bool),
		edges: make(map[*Node][]*edge),
	}
}

func (g *UndirectedGraph) AddNode(n *Node) error {
	if n == nil {
		return errors.New("node is nil")
	}

	if _, ok := g.nodes[n]; ok {
		return errors.New("node already added")
	}

	g.nodes[n] = true
	return nil
}

// AddEdge on undirected graphs assumes adding n1 -> n2 edge, as well as n2 -> n1 edge with identical weights
func (g *UndirectedGraph) AddEdge(n1, n2 *Node, weight int) error {
	if !g.nodes[n1] || !g.nodes[n2] {
		return errors.New("the provided node(s) are not added to the graph")
	}

	if _, ok := g.edges[n1]; !ok {
		g.edges[n1] = make([]*edge, 0)
	}

	if _, ok := g.edges[n2]; !ok {
		g.edges[n2] = make([]*edge, 0)
	}

	g.edges[n1] = append(g.edges[n1], &edge{N1: n1, N2: n2, Weight: weight})
	g.edges[n2] = append(g.edges[n2], &edge{N2: n1, N1: n2, Weight: weight})

	return nil
}

func (g *UndirectedGraph) Edges(n *Node) []*edge {
	return g.edges[n]
}

func (g *UndirectedGraph) NodeExists(n *Node) bool {
	return g.nodes[n]
}
