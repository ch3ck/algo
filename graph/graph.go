package graph

import "fmt"

/*import "os"
import "strconv"
import "fmt"*/

//Heapsort for integers in Golang
//based on Professor Mark Allen Weiss' solution in C
//http://users.cs.fiu.edu/~weiss/

// Node
type ID string

type Node interface {
	Id() ID
	Edges() []ID
}

type node struct {
	id    ID
	edges []ID
}

func (n node) Id() ID {
	return n.id
}

func (n node) Edges() []ID {
	return n.edges
}

// Graph
type Graph interface {
	AddNode(id ID, edges []ID) Node
	BreadthFirstTraversal(start ID, visitor NodeVisitorFunc) error
}

// returning an error terminates the traversal
type NodeVisitorFunc func(n Node, depth int) error

type graph struct {
	nodes map[ID]*node
}

func NewDirectedGraph() Graph {
	return &graph{
		nodes: map[ID]*node{},
	}
}

func (g graph) AddNode(id ID, edges []ID) Node {
	node := &node{
		id:    id,
		edges: edges,
	}
	g.nodes[id] = node
	return node
}

// Visits each node in breadth-first order. if the visitor function returns and error
// traversal is terminated and the error returned by this function.
func (g graph) BreadthFirstTraversal(start ID, visitor NodeVisitorFunc) error {
	toVisitIds := []ID{}
	visitedIds := map[ID]interface{}{}

	depth := 0
	toVisitIds = append(toVisitIds, start)

	for len(toVisitIds) != 0 {
		currId := toVisitIds[0]
		visitedIds[currId] = nil

		if node := g.nodes[currId]; node != nil {
			if err := visitor(node, depth); err != nil {
				return err
			}
			for _, id := range node.edges {
				if _, ok := visitedIds[id]; !ok {
					toVisitIds = append(toVisitIds, id)
				}
			}
		} else {
			return fmt.Errorf("inconsistent graph, missing node with id %s", currId)
		}

		toVisitIds = toVisitIds[1:]
	}

	return nil
}
