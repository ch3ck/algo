package graph_test

import (
	"fmt"
	"github.com/Ch3ck/AlGo/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreadthFirstTraversal(t *testing.T) {
	t.Run("directed graph depth first traversal", func(t *testing.T) {
		// create data
		g := graph.NewDirectedGraph()
		g.AddNode("0", []graph.ID{"1a", "1b", "1c"})
		g.AddNode("1a", []graph.ID{"2a1"})
		g.AddNode("1b", []graph.ID{"2b1", "2b2"})
		g.AddNode("1c", []graph.ID{"2c1"})
		g.AddNode("2a1", []graph.ID{"3a1"})
		g.AddNode("2b1", []graph.ID{"0"})
		g.AddNode("2b2", []graph.ID{"1a"})
		g.AddNode("2c1", []graph.ID{"2c1"})
		g.AddNode("3a1", []graph.ID{})

		fmt.Printf("input: %+v\n", g)

		expected := []graph.ID{"0", "1a", "1b", "1c", "2a1", "2b1", "2b2", "2c1", "3a1"}
		actual := []graph.ID{}
		err := g.BreadthFirstTraversal("0", func(n graph.Node, depth int) error {
			actual = append(actual, n.Id())
			return nil
		})

		assert.NoError(t, err, "traversal error")
		assert.Equal(t, expected, actual, "traversal order failed")
	})
}
