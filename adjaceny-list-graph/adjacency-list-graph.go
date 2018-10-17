// Copyright 2018 Anthony A. Abeo. All rights reserved.
// Use of this source code is governed by a `MIT License`.
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

import (
	"github.com/Ch3ck/AlGo/queuestack"
)

// A vertex (or node) in the graph.
// Each vertex has the following attributes:
//
// id (int): a number to help access the node in O(1) time
//	item (interface): the item stored in the vertex
//	outgoingEdges (list.List): a list of the vertices, that along
// 							   with the current vertex form an outgoing edge.
//	incomingEdges (list.List): a list of the vertices, that along
//							   with the current vertex form an outgoing edge.
type vertex struct {
	id            int
	item          interface{}
	outgoingEdges []*edge
	incomingEdges []*edge
}

// An edge in the directed graph.
// Each edge is represented as a 2-tuple struct that has the following attributes
//
// weight (interface): a value denoted the weight of an edge (when dealing with weighted graphs).
//						It is set to nil for unweighted graphs.
// fromVertex (vertex): the vertex this edge starts from. The outgoing vertex.
// toVertex (vertex): the vertex this edge is incident on. The incoming vertex.
type edge struct {
	id         int
	weight     interface{}
	fromVertex *vertex
	toVertex   *vertex
}

func NewEdge() *edge {
	return &edge{
		0,
		"sugar",
		new(vertex),
		new(vertex),
	}
}

// The Directed Adjacency List Graph.
//
// ATTRIBUTES
//
// primaryStructure (map[int]vertex): holds the vertices in the graph. A map is used
//                                    because it provides O(1) access to a node.
// numVertices (int): The total number of vertices in the graph.
// numEdges (int): The total number of edges in the graph.
type DirectedAdjacencyListGraph struct {
	primaryStructure map[int]*vertex
	numVertices      int
	numEdges         int
}

func NewDALGraph() *DirectedAdjacencyListGraph {
	return &DirectedAdjacencyListGraph{
		make(map[int]*vertex),
		0,
		0,
	}
}

// NumVertices returns the total number of vertices currently in the graph.
func (dalg *DirectedAdjacencyListGraph) NumVertices() int {
	return dalg.numVertices
}

// NumEdges returns the total number of edges currently in the graph.
func (dalg *DirectedAdjacencyListGraph) NumEdges() int {
	return dalg.numEdges
}

// Edges returns a list of all the edges in the graph.
func (dalg *DirectedAdjacencyListGraph) Edges() []*edge {
	var edges = make([]*edge, dalg.numEdges)
	for _, v := range dalg.primaryStructure {
		edges = append(edges, v.outgoingEdges...)
	}

	return edges
}

// Edges returns a list of all the edges in the graph
func (dalg *DirectedAdjacencyListGraph) Vertices() []*vertex {
	var vertices = make([]*vertex, dalg.numVertices)
	for _, v := range dalg.primaryStructure {
		vertices = append(vertices, v)
	}

	return vertices
}

// GetEdge returns the `edge` that start from `fromVertex` and ends at `toVertex`
// if one exists; otherwise return null.
//
// ARGUMENTS
//
// fromVertex (vertex): the vertex this edge starts from. The outgoing vertex.
// toVertex (vertex): the vertex this edge is incident on. The incoming vertex.
func (dalg *DirectedAdjacencyListGraph) GetEdge(fromVertexID, toVertexID int) *edge {
	fromVertex := dalg.primaryStructure[fromVertexID]
	toVertex := dalg.primaryStructure[toVertexID]
	if fromVertex != nil && toVertex != nil {
		for _, e := range fromVertex.outgoingEdges {
			if e.toVertex == toVertex {
				return e
			}
		}
	}

	return nil
}

// OutDegree returns the number of outgoing edges from vertex with the specified ID.
//
// ARGUMENTS
//
// id (int): ID of the vertex of interest
func (dalg *DirectedAdjacencyListGraph) OutDegree(id int) int {
	if dalg.primaryStructure[id] != nil {
		return len(dalg.primaryStructure[id].outgoingEdges)
	}

	return -1
}

// InDegree returns the number of outgoing edges from vertex with the specified ID.
//
// ARGUMENTS
//
// id (int): ID of the vertex of interest
func (dalg *DirectedAdjacencyListGraph) InDegree(id int) int {
	if dalg.primaryStructure[id] != nil {
		return len(dalg.primaryStructure[id].incomingEdges)
	}

	return -1
}

// OutgoingEdges Returns a list of all outgoing edges from vertex with the specified ID.
//
// ARGUMENTS
//
// id (int): ID of the vertex of interest
func (dalg *DirectedAdjacencyListGraph) OutgoingEdges(id int) []*edge {
	if dalg.primaryStructure[id] != nil {
		return dalg.primaryStructure[id].outgoingEdges
	}

	return nil
}

// IncomingEdges Returns a list of all outgoing edges from vertex with the specified ID.
//
// ARGUMENTS
//
// id (int): ID of the vertex of interest
func (dalg *DirectedAdjacencyListGraph) IncomingEdges(id int) []*edge {
	if dalg.primaryStructure[id] != nil {
		return dalg.primaryStructure[id].outgoingEdges
	}

	return nil
}

// InsertVertex Creates and returns a new Vertex storing element `item`.
//
// ARGUMENTS
// item (interface): the item the vertex will hold.
func (dalg *DirectedAdjacencyListGraph) InsertVertex(item interface{}) *vertex {
	if item != nil {
		v := &vertex{
			dalg.numVertices,
			item,
			make([]*edge, 0),
			make([]*edge, 0),
		}
		dalg.primaryStructure[dalg.numVertices] = v
		dalg.numVertices++

		return v
	}

	return nil
}

// InsertEdge Creates and returns a new Edge from vertex `fromVertex` to vertex `toVertex`,
// storing element x; an error occurs if there already exists an edge from `fromVertex` to `toVertex`
//
// ARGUMENTS
//
// fromVertex (vertex): the vertex this edge starts from. The outgoing vertex.
// toVertex (vertex): the vertex this edge is incident on. The incoming vertex.
// item (interface): the item the edge will hold (in the case of a weighted graph).
func (dalg *DirectedAdjacencyListGraph) InsertEdge(fromVertexID, toVertexID int, item interface{}) interface{} {
	if dalg.GetEdge(fromVertexID, toVertexID) == nil {
		fromVertex := dalg.primaryStructure[fromVertexID]
		toVertex := dalg.primaryStructure[toVertexID]
		if fromVertex != nil && toVertex != nil {
			e := &edge{
				dalg.numEdges,
				item,
				fromVertex,
				toVertex,
			}

			fromVertex.outgoingEdges = append(fromVertex.outgoingEdges, e)
			toVertex.incomingEdges = append(toVertex.incomingEdges, e)

			dalg.numEdges++

			return e
		}
	}

	return nil
}

// RemoveVertex Removes vertex with the specified ID and all its incident edges from the graph.
//
// ARGUMENTS
//
// id (int): ID of the vertex of interest
func (dalg *DirectedAdjacencyListGraph) RemoveVertex(id int) bool {
	if dalg.primaryStructure[id] != nil {
		delete(dalg.primaryStructure, id)
		dalg.numVertices--

		return true
	}

	return false
}

// RemoveEdge removes edge e from the graph.
//
// ARGUMENTS
//
// e (edge): the edge to be removed
func (dalg *DirectedAdjacencyListGraph) RemoveEdge(e *edge) bool {
	if e != nil && e.id >= 0 && e.id < len(e.fromVertex.outgoingEdges) {

		e.fromVertex.outgoingEdges[e.id] = e.fromVertex.outgoingEdges[len(e.fromVertex.outgoingEdges)-1]
		e.fromVertex.outgoingEdges = e.fromVertex.outgoingEdges[:len(e.fromVertex.outgoingEdges)-1]

		e.toVertex.incomingEdges[e.id] = e.toVertex.incomingEdges[len(e.toVertex.incomingEdges)-1]
		e.toVertex.incomingEdges = e.toVertex.incomingEdges[:len(e.toVertex.incomingEdges)-1]

		dalg.numEdges--

		return true
	}

	return false
}

func (dalg *DirectedAdjacencyListGraph) BFS() []int {
	var l []int
	visitedTable := make(map[int]bool, dalg.numVertices)
	for key := range dalg.primaryStructure {
		visitedTable[key] = false
	}

	for key, vertex := range dalg.primaryStructure {
		if !visitedTable[key] {
			l = bfs(vertex, visitedTable)
		}
	}

	return l
}

func bfs(v *vertex, visitedTable map[int]bool) []int {
	l := make([]int, 0)
	queue := queuestack.NewQueue()
	queue.Enqueue(v)
	visitedTable[v.id] = true

	l = append(l, v.id)

	for !queue.Empty() {
		if node, err := queue.Dequeue(); err == nil {
			for _, ed := range node.(*vertex).outgoingEdges {
				if !visitedTable[ed.toVertex.id] {
					visitedTable[ed.toVertex.id] = true
					queue.Enqueue(ed.toVertex)

					l = append(l, ed.toVertex.id)
				}
			}
		}
	}

	return l
}