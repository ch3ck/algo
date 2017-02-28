package kruskal

// Type holding parent pointer and rank for a DSU node
type Node struct {
	parent *Node
	rank   int
}

// Recursive implementation of find with path compression
func Find(node *Node) *Node {
	// If a node is root of it's component then it has it's parent pointer set to nil
	if node.parent == nil {
		return node
	} else {
		node.parent = Find(node.parent)
		return node.parent
	}
}

// Implementation of union with ranks, returns a node being the root of the new component
func Union(a, b *Node) *Node {
	a = Find(a)
	b = Find(b)

	if a == b {
		return a
	}

	// Swap nodes if needed so that b has higher rank
	if a.rank < b.rank {
		a, b = b, a
	}

	b.parent = a
	if a.rank == b.rank {
		a.rank++
	}

	return a
}
