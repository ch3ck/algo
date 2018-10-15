package adjaceny_list_graph

import "testing"

// Setup
var dalg = NewDALGraph()

func TestDirectedAdjacencyListGraph_NumEdges(t *testing.T) {
	if dalg.numEdges != 0 {
		t.Errorf("Number of edges for newly created graph was incorrect, got: %d, want: %d.", dalg.numEdges, 0)
	}
}

func TestDirectedAdjacencyListGraph_NumVertices(t *testing.T) {
	if dalg.numVertices != 0 {
		t.Errorf("Incorrect, got: %d, want: %d.", dalg.numVertices, 0)
	}
}
