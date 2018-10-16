package adjaceny_list_graph

import (
	"testing"
)

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

	if dalg.primaryStructure[0].outgoingEdges[0].weight != "brother" ||
		dalg.primaryStructure[1].incomingEdges[0].weight != "brother" {

		t.Error("edge between vertex 0 and 1 broken")
	}

	e = dalg.InsertEdge(1, 4, "brother")
	if e != nil {
		t.Error("Incorrect, should be nil")
	}

	e = dalg.InsertEdge(0, 1, "brother")
	if e != nil {
		t.Error("Incorrect, should be nil. Edge already exist")
	}
}

func TestDirectedAdjacencyListGraph_GetEdge(t *testing.T) {
	e := dalg.GetEdge(0, 1)
	if e == nil {
		t.Error("Incorrect. Edge actually exist")
	}

	if e.weight != "brother" {
		t.Errorf("Incorrect, got: %d, expected: %v.", e.weight, "brother")
	}

	e = dalg.GetEdge(1, 4)
	if e != nil {
		t.Error("Incorrect. Edge DOES NOT actually exist. e should be nil")
	}
}

func TestDirectedAdjacencyListGraph_RemoveEdge(t *testing.T) {
	e := dalg.GetEdge(0, 1)
	if e == nil {
		t.Error("Incorrect. Edge actually exist")
	}

	dalg.RemoveEdge(e)

	e = dalg.GetEdge(0, 1)
	if e != nil {
		t.Error("Incorrect. Edge does not exist. Should be nil")
	}

	if dalg.numEdges != 0 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.numEdges, 0)
	}

	e = dalg.GetEdge(1, 4)
	if e != nil {
		t.Error("Incorrect. Edge DOES NOT actually exist. e should be nil")
	}

	if dalg.RemoveEdge(e) {
		t.Error("Incorrect. Edge DOES NOT actually exist.")
	}

	if dalg.RemoveEdge(NewEdge()) {
		t.Error("Incorrect. Edge DOES NOT actually exist.")
	}
}

func TestDirectedAdjacencyListGraph_RemoveVertex(t *testing.T) {
	dalg.RemoveVertex(0)
	if dalg.numVertices != 1 {
		t.Errorf("Incorrect, got: %d, expected: %v.", dalg.numVertices, 1)
	}

	if dalg.RemoveVertex(0) {
		t.Error("Incorrect. Vertex DOES NOT actually exist.")
	}
}
