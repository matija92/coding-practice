package graph

import "errors"

type UnionFindGraph struct {
	root  map[*Node]*Node
	edges map[*Node][]*edge // For reference
}

func NewUnionFind() *UnionFindGraph {
	return &UnionFindGraph{
		root:  make(map[*Node]*Node),
		edges: make(map[*Node][]*edge),
	}
}

func (g *UnionFindGraph) AddNode(n *Node) error {
	if n == nil {
		return ErrInvalidInput
	}

	if g.NodeExists(n) {
		return ErrNodeAlreadyAdded
	}

	g.root[n] = n
	return nil
}

func (g *UnionFindGraph) AddEdge(n1, n2 *Node, weight int) error {
	if !g.NodeExists(n1) || !g.NodeExists(n2) {
		return errors.New("the provided node(s) are not added to the graph")
	}

	// Record edge
	if _, ok := g.edges[n1]; !ok {
		g.edges[n1] = make([]*edge, 0)
	}
	g.edges[n1] = append(g.edges[n1], &edge{
		N1:     n1,
		N2:     n2,
		Weight: weight,
	})

	return g.Union(n1, n2)
}

func (g *UnionFindGraph) Edges(n *Node) []*edge {
	return g.edges[n]
}

func (g *UnionFindGraph) NodeExists(n *Node) bool {
	return g.root[n] != nil
}

// Find returns the root node of the given node
func (g *UnionFindGraph) Find(n *Node) *Node {
	return g.root[n]
}

// Union adds an edge between 2 nodes
func (g *UnionFindGraph) Union(n1, n2 *Node) error {
	if !g.NodeExists(n1) || !g.NodeExists(n2) {
		return ErrNodeNotInGraph
	}

	root1 := g.Find(n1)
	root2 := g.Find(n2)

	if root1 != root2 {
		for node, root := range g.root {
			if root == root2 {
				g.root[node] = root1
			}
		}
	}
	return nil
}
