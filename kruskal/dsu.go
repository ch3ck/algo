package kruskal

type Node struct {
	parent *Node
	rank   int
}

func Find(node *Node) *Node {
	if node.parent == nil {
		return node
	} else {
		node.parent = Find(node.parent)
		return node.parent
	}
}

func Union(a, b *Node) *Node {
	a = Find(a)
	b = Find(b)

	if a == b {
		return a
	}

	if a.rank < b.rank {
		a, b = b, a
	}

	b.parent = a
	if a.rank == b.rank {
		a.rank++
	}

	return a
}
