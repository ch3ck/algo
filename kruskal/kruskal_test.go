package kruskal

import "testing"

func TestKruskal(t *testing.T) {
	edges := []Edge{{4, 5, 10}}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			edges = append(edges, Edge{i, j, j - i})
		}
	}

	forestEdges := MinimumSpanningForest(6, edges)

	if len(forestEdges) != 4 {
		t.Error("Returned set of edges has invalid size")
	}

	forestEdgesSet := make(map[Edge]bool)
	for _, edge := range forestEdges {
		forestEdgesSet[edge] = true
	}

	answer := []Edge{{4, 5, 10}}

	for i := 0; i < 3; i++ {
		answer = append(answer, Edge{i, i + 1, 1})
	}

	for _, edge := range answer {
		if _, ok := forestEdgesSet[edge]; !ok {
			t.Error("Returned set of edges is not a minimum spanning forest")
		}
	}
}
