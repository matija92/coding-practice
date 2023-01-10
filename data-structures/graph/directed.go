package graph

type DirectedGraph struct {
	nodes map[*Node]bool // Lookup map for nodes
	edges map[*Node][]*Edge
}

func NewDirected() *DirectedGraph {
	return &DirectedGraph{
		nodes: make(map[*Node]bool),
		edges: make(map[*Node][]*Edge),
	}
}

func (g *DirectedGraph) AddNode(n *Node) error {
	if n == nil {
		return ErrInvalidInput
	}

	if _, ok := g.nodes[n]; ok {
		return ErrNodeAlreadyAdded
	}

	g.nodes[n] = true
	return nil
}

// AddEdge on directed graphs assumes n1 -> n2 as a single edge
func (g *DirectedGraph) AddEdge(n1, n2 *Node, weight int) error {
	if !g.nodes[n1] || !g.nodes[n2] {
		return ErrNodeNotInGraph
	}

	if _, ok := g.edges[n1]; !ok {
		g.edges[n1] = make([]*Edge, 0)
	}

	g.edges[n1] = append(g.edges[n1], &Edge{N1: n1, N2: n2, Weight: weight})

	return nil
}

func (g *DirectedGraph) Edges(n *Node) []*Edge {
	return g.edges[n]

}

func (g *DirectedGraph) Nodes() []*Node {
	nodes := make([]*Node, len(g.nodes))
	i := 0
	for n := range g.nodes {
		nodes[i] = n
		i++
	}
	return nodes
}

func (g *DirectedGraph) NodeExists(n *Node) bool {
	return g.nodes[n]
}
