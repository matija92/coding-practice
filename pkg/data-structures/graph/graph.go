package graph

type Graph interface {
	AddNode(n *Node) error
	AddEdge(n1, n2 *Node, weight int) error
	Edges(n *Node) []*edge
	NodeExists(n *Node) bool
}

type Node struct {
	name string
	val  int
}

func NewNode(name string, value int) *Node {
	return &Node{
		name: name,
		val:  value,
	}
}

type edge struct {
	N1     *Node // start node
	N2     *Node // end node
	Weight int
}
