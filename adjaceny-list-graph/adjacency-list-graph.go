// Copyright 2018 alGO. All rights reserved.
// Use of this source code is governed by the `MIT License`.
// license that can be found in the LICENSE file.

// SUMMARY
//
// This module implements a graph data structure using the adjacency
// list implementation as defined by Goodrich, Tamassia & Goldwasser (2014)
// in [1]. This implementation however makes some changes; most notably
// is the use of a map data structure for the `primary structure` that holds
// the vertices as compared to the array based data structure the authors suggested.
// The Graph Abstract Data Type implemented is the same as the the one proposed
// by the authors.
//
// As suggested by the authors, each vertex in the graph uses an will have
// both outgoing and incoming list of edges, since the graph is directed.
//
//
// REFERENCE
//
// [1] Goodrich, M. T., Tamassia, R., & Goldwasser, M. H. (2014).
// 	   Data structures and algorithms in Java. John Wiley & Sons.

package adjaceny_list_graph

import "container/list"

// A vertex (or node) in the graph.
// Each vertex has the following attributes:
//
// 		id (int): a number to help access the node in O(1) time
//		item (interface): the item stored in the vertex
//		outgoingEdges (list.List): a list of the vertices, that along
// 								   with the current vertex form an outgoing edge.
//		incomingEdges (list.List): a list of the vertices, that along
//								   with the current vertex form an outgoing edge.
type vertex struct {
	id            int
	item          interface{}
	outgoingEdges list.List
	incomingEdges list.List
}

type edge struct {
	fromVertex vertex
	toVertex   vertex
}

type DirectedAdjacencyListGraph struct {
	primaryStructure map[int]vertex
}

func (dalg *DirectedAdjacencyListGraph) NumVertices() int {
	return 0
}

func (dalg *DirectedAdjacencyListGraph) NumEdges() int {
	return 0
}

func (dalg *DirectedAdjacencyListGraph) Edges() int {
	return 0
}

func (dalg *DirectedAdjacencyListGraph) GetEdge(fromVertex, toVertex vertex) edge {
	return edge{
		vertex{},
		vertex{},
	}
}

func (dalg *DirectedAdjacencyListGraph) OutDegree() int {
	return 0
}

func (dalg *DirectedAdjacencyListGraph) InDegree() int {
	return 0
}

func (dalg *DirectedAdjacencyListGraph) OutgoingEdges(v vertex) list.List {
	return v.outgoingEdges
}

func (dalg *DirectedAdjacencyListGraph) IncomingEdges(v vertex) list.List {
	return v.outgoingEdges
}

func (dalg *DirectedAdjacencyListGraph) InsertVertex(item interface{}) {}

func (dalg *DirectedAdjacencyListGraph) InsertEdge(fromVertex, toVertex vertex, item interface{}) {}

func (dalg *DirectedAdjacencyListGraph) RemoveVertex(v vertex) {}

func (dalg *DirectedAdjacencyListGraph) RemoveEdge() {}
