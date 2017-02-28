package kruskal

import "testing"

// Tests the DSU implementation
func TestDsu(t *testing.T) {
	// Creates 5 one-element sets
	nodes := make([]Node, 5)

	// Performs find on an one-element set
	if Find(&nodes[0]) != &nodes[0] {
		t.Error("Find on an newly created element should return itself")
	}

	Union(&nodes[0], &nodes[1])
	Union(&nodes[2], &nodes[3])

	if Find(&nodes[0]) != Find(&nodes[1]) {
		t.Error("Nodes after union should have equal representatives")
	}

	// Unions nodes 0, 1, 2, 3 into one set
	Union(&nodes[0], &nodes[3])
	Union(&nodes[1], &nodes[3])

	if Find(&nodes[0]) == Find(&nodes[4]) {
		t.Error("Nodes in different sets should have different representatives")
	}

	if Find(&nodes[0]) != Find(&nodes[2]) {
		t.Error("Nodes in the same set should have equal representatives")
	}
}
