// Package provides an implementation of the Kruskal's minimum spanning forest algorithm
// together with an implementation of disjoint set data structure
package kruskal

import "sort"

// Type for edge data - endpoints u, v and length
type Edge struct {
	u, v, length int
}

// Type required for calling the sort.Sort method
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

// Computes the minimum spanning forest in a graph with vertices numbered from 0 to n-1
// and a given set of edges.
// Returns the subset of edges forming the minimum spanning forest.
func MinimumSpanningForest(n int, edges []Edge) []Edge {
	forestEdges := make([]Edge, 0)

	// Disjoint set union that holds the division of the graph into connected components
	nodes := make([]Node, n)

	// Sort edges increasingly by length using the Less() function
	sort.Sort(Edges(edges))

	for _, edge := range edges {
		uNode := &nodes[edge.u]
		vNode := &nodes[edge.v]

		// If an edge connects two different components, add it to the forest
		if Find(uNode) != Find(vNode) {
			Union(uNode, vNode)
			forestEdges = append(forestEdges, edge)
		}
	}

	return forestEdges
}
