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

func TestIntegration(t *testing.T)  {
	// Create a new graph
	g := NewDALGraph()

	// Insert all the vertices
	v1 := g.InsertVertex(1)
	v2 := g.InsertVertex(2)
	v3 := g.InsertVertex(3)
	v4 := g.InsertVertex(4)
	v5 := g.InsertVertex(5)

	// Insert all the edges
	e12 := g.InsertEdge(v1.id, v2.id, "12")
	g.InsertEdge(v1.id, v4.id, "14")
	g.InsertEdge(v3.id, v1.id, "31")
	e34 := g.InsertEdge(v3.id, v4.id, "34")
	g.InsertEdge(v4.id, v5.id, "45")
	g.InsertEdge(v5.id, v3.id, "53")

	e912 := g.InsertEdge(9, 12, "912")
	if e912 != nil {t.Errorf("Incorrect, cannot insert such edge")}

	// Checking for the number of edges and vertices
	if g.NumVertices() != 5 {t.Errorf("Incorrect, got: %d, expected: %v.", g.NumVertices(), 5)}
	if g.NumEdges() != 6 {t.Errorf("Incorrect, got: %d, expected: %v.", g.NumEdges(), 6)}

	// Testing GetEdge()
	e := g.GetEdge(v1.id, v2.id)
	if e.id != e12.(*edge).id {t.Errorf("Incorrect, got: %d, expected: %v.", e.id, e12.(*edge).id)}
	if e.weight != e12.(*edge).weight {t.Errorf("Incorrect, got: %d, expected: %v.", e.weight, e12.(*edge).weight)}

	e = g.GetEdge(v3.id, v4.id)
	if e.id != e34.(*edge).id {t.Errorf("Incorrect, got: %d, expected: %v.", e.id, e34.(*edge).id)}
	if e.weight != e34.(*edge).weight {t.Errorf("Incorrect, got: %d, expected: %v.", e.weight, e34.(*edge).weight)}

	e = g.GetEdge(5, 3)
	if e != nil {t.Error("Incorrect. That edge is invalid")}

	// Testing OutDegree
	if g.OutDegree(v1.id) != 2 {t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(v1.id), 2)}
	if g.OutDegree(v3.id) != 2 {t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(v3.id), 2)}
	if g.OutDegree(v5.id) != 1 {t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(v5.id), 1)}
	if g.OutDegree(9) != -1 {t.Errorf("Incorrect, got: %d, expected: %v.", g.OutDegree(9), -1)}

	// Testing InDegree
	if g.InDegree(v4.id) != 2 {t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(v4.id), 2)}
	if g.InDegree(v3.id) != 1 {t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(v3.id), 1)}
	if g.InDegree(v2.id) != 1 {t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(v2.id), 1)}
	if g.InDegree(9) != -1 {t.Errorf("Incorrect, got: %d, expected: %v.", g.InDegree(9), -1)}
}