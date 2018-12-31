// This package provide a algorithm for distance of all pairs problems
// It solves this problem with Floyd-Warshall algorithm in O(N^3)
package allpairs

// Max Integer
const INF = int(^uint(0) >> 2)

// Type for edge data - endpoints u, v and length
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

// Computes the distance of all pairs
func allPairsDistance(n int, edges []Edge) [][]int {
	distance := make([][]int, n)
	for i := range distance {
		distance[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			distance[i][j] = INF
		}
	}

	for _, edge := range edges {
		if distance[edge.u][edge.v] > edge.length {
			distance[edge.u][edge.v] = edge.length
		}
	}

	for i := 0; i < n; i++ {
		distance[i][i] = 0
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distance[i][j] > distance[i][k]+distance[k][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
				}
			}
		}
	}

	return distance
}
