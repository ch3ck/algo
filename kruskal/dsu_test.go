package kruskal

import "testing"

func TestDsu(t *testing.T) {
	nodes := make([]Node, 5)

	if Find(&nodes[0]) != &nodes[0] {
		t.Error("Find on an newly created element should return itself")
	}

	Union(&nodes[0], &nodes[1])
	Union(&nodes[2], &nodes[3])

	if Find(&nodes[0]) != Find(&nodes[1]) {
		t.Error("Nodes after union should have equal representatives")
	}

	Union(&nodes[0], &nodes[3])
	Union(&nodes[1], &nodes[3])

	if Find(&nodes[0]) == Find(&nodes[4]) {
		t.Error("Nodes in different sets should have different representatives")
	}

	if Find(&nodes[0]) != Find(&nodes[2]) {
		t.Error("Nodes in the same set should have equal representatives")
	}
}
