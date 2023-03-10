package graph

import "errors"

type UndirectedGraph struct {
	nodes map[*Node]bool // Lookup map for nodes
	edges map[*Node][]*Edge
}

func NewUndirected() *UndirectedGraph {
	return &UndirectedGraph{
		nodes: make(map[*Node]bool),
		edges: make(map[*Node][]*Edge),
	}
}

func (g *UndirectedGraph) AddNode(n *Node) error {
	if n == nil {
		return ErrInvalidInput
	}

	if _, ok := g.nodes[n]; ok {
		return ErrNodeAlreadyAdded
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
		g.edges[n1] = make([]*Edge, 0)
	}

	if _, ok := g.edges[n2]; !ok {
		g.edges[n2] = make([]*Edge, 0)
	}

	for _, edge := range g.edges[n1] {
		if edge.N2 == n2 {
			return ErrEdgeAlreadyAdded
		}
	}

	for _, edge := range g.edges[n2] {
		if edge.N2 == n1 {
			return ErrEdgeAlreadyAdded
		}
	}

	g.edges[n1] = append(g.edges[n1], &Edge{N1: n1, N2: n2, Weight: weight})
	g.edges[n2] = append(g.edges[n2], &Edge{N2: n1, N1: n2, Weight: weight})

	return nil
}

func (g *UndirectedGraph) Edges(n *Node) []*Edge {
	return g.edges[n]
}

func (g *UndirectedGraph) Nodes() []*Node {
	nodes := make([]*Node, len(g.nodes))
	i := 0
	for n := range g.nodes {
		nodes[i] = n
		i++
	}
	return nodes
}

func (g *UndirectedGraph) NodeExists(n *Node) bool {
	return g.nodes[n]
}
