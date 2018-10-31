package allpairs

import "testing"

// Tests the Floyd Warshall all pairs algorithm
func TestAllPairs(t *testing.T) {
	edges := []Edge{{4, 5, 10}}

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			edges = append(edges, Edge{i, j, j - i})
		}
	}

	// For the given test the answer has cost 13 and is unique

	distance := allPairsDistance(6, edges)

	if len(distance) != 6 {
		t.Error("Returned array of distance have wrong size")
	}

	if distance[4][5] != 10 {
		t.Error("Distance from 4 to 5 is wrong")
	}

	if distance[1][5] != INF {
		t.Error("Distance from 1 to 5 is wrong")
	}

	if distance[1][2] != 1 {
		t.Error("Distance from 1 to 2 is wrong")
	}

	if distance[3][3] != 0 {
		t.Error("Distance from 3 to 3 is wrong")
	}

	if distance[3][5] != INF {
		t.Error("Distance from 3 to 4 is wrong")
	}
}
