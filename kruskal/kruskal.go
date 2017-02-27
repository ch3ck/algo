package kruskal

import "sort"

type Edge struct {
	u, v, length int
}

type Edges []Edge

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].length < edges[j].length
}

func MinimumSpanningForest(n int, edges []Edge) []Edge {
	forestEdges := make([]Edge, 0)

	nodes := make([]Node, n)

	sort.Sort(Edges(edges))

	for _, edge := range edges {
		uNode := &nodes[edge.u]
		vNode := &nodes[edge.v]

		if Find(uNode) != Find(vNode) {
			Union(uNode, vNode)
			forestEdges = append(forestEdges, edge)
		}
	}

	return forestEdges
}
