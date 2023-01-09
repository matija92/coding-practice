package graph

type Graph interface {
	AddNode(n *Node) error
	AddEdge(n1, n2 *Node, weight int) error
	Edges(n *Node) []*edge
	Nodes() []*Node
	NodeExists(n *Node) bool
}

type graphError string

func (e graphError) Error() string {
	return string(e)
}

const (
	ErrInvalidInput     graphError = "input is invalid"
	ErrNodeNotInGraph   graphError = "node not added"
	ErrNodeAlreadyAdded graphError = "node already added"
	ErrEdgeAlreadyAdded graphError = "edge already added"
)

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

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Value() int {
	return n.val
}

func (n *Node) SetValue(x int) {
	n.val = x
}

type edge struct {
	N1     *Node // start node
	N2     *Node // end node
	Weight int
}
