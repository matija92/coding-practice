package islands

import (
	"coding-practice/data-structures/graph"
)

// Problem:
// Given a graph, count the number of islands
func NoOfIslands(g *graph.UnionFindGraph) int {
	islandsMap := make(map[*graph.Node]interface{})
	for _, node := range g.Nodes() {
		islandsMap[g.Find(node)] = true
	}
	return len(islandsMap)
}
