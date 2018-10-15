package adjaceny_list_graph

import "testing"

// Setup
var dalg = NewDALGraph()

func TestDirectedAdjacencyListGraph_NumEdges(t *testing.T) {
	if dalg.numEdges != 0 {
		t.Errorf("Incorrect, got: %d, expected: %d.", dalg.numEdges, 0)
	}
}

func TestDirectedAdjacencyListGraph_NumVertices(t *testing.T) {
	if dalg.numVertices != 0 {
		t.Errorf("Incorrect, got: %d, expected: %d.", dalg.numVertices, 0)
	}
}

func TestDirectedAdjacencyListGraph_InsertVertex(t *testing.T) {
	v := dalg.InsertVertex("thor")
	if v.item != "thor" {
		t.Errorf("Incorrect, got: %d, expected: %s.", v.item, "thor")
	}

	v = dalg.InsertVertex("loki")
	if v.item != "loki" {
		t.Errorf("Incorrect, got: %d, expected: %s.", v.item, "loki")
	}

	v = dalg.InsertVertex(nil)
	if v != nil {
		t.Errorf("Incorrect, got: %d, expected: %v.", v, nil)
	}

	if dalg.numVertices != 2 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.numVertices, 1)
	}
}

func TestDirectedAdjacencyListGraph_InsertEdge(t *testing.T) {
	e := dalg.InsertEdge(0, 1, "brother")
	if e.(*edge).weight != "brother" {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.(edge).weight, "brother")
	}

	if dalg.numEdges != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.numEdges, 1)
	}
}