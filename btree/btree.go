package btree

//Index interface of data container of Node
type Index interface {
	LessThan(Index) bool
	EqualsTo(Index) bool
}

//Node Tree element struct
type Node struct {
	Data  Index
	Score int
	Edges [2]*Node
}

// Public

// Insert a node into the AVL tree.
func Insert(tree **Node, data Index) {
	*tree, _ = insertInRightOf(*tree, data)
}

// Remove a single item from an AVL tree.
func Remove(tree **Node, data Index) {
	*tree, _ = removeFromRightOf(*tree, data)
}

// Private

func opp(d int) int {
	return 1 - d
}

//singleRotation
func singleRotation(parent *Node, direction int) (b *Node) {
	b = parent.Edges[opp(direction)]
	parent.Edges[opp(direction)] = b.Edges[direction]
	b.Edges[direction] = parent
	return
}

//doubleRotation
func doubleRotation(parent *Node, d int) (b *Node) {
	b = parent.Edges[opp(d)].Edges[d]
	parent.Edges[opp(d)].Edges[d] = b.Edges[opp(d)]
	b.Edges[opp(d)] = parent.Edges[opp(d)]
	parent.Edges[opp(d)] = b

	b = parent.Edges[opp(d)]
	parent.Edges[opp(d)] = b.Edges[d]
	b.Edges[d] = parent
	return
}

func adjust(parent *Node, dir, bal int) {
	n := parent.Edges[dir]
	nn := n.Edges[opp(dir)]
	switch nn.Score {
	case 0:
		parent.Score = 0
		n.Score = 0
	case bal:
		parent.Score = bal
		n.Score = 0
	default:
		parent.Score = 0
		n.Score = bal
	}
	nn.Score = 0
}

func insertScoreBalance(node *Node, dir int) *Node {
	n := node.Edges[dir]
	bal := 2*dir - 1
	if n.Score == bal {
		node.Score = 0
		n.Score = 0
		return singleRotation(node, opp(dir))
	}
	adjust(node, dir, bal)
	return doubleRotation(node, opp(dir))
}

//insertInRightOf
func insertInRightOf(node *Node, data Index) (*Node, bool) {
	if node == nil {
		return &Node{Data: data}, false
	}
	dir := 0
	if node.Data.LessThan(data) {
		dir = 1
	}
	var done bool
	node.Edges[dir], done = insertInRightOf(node.Edges[dir], data)
	if done {
		return node, true
	}
	node.Score += 2*dir - 1
	switch node.Score {
	case 0:
		return node, true
	case 1, -1:
		return node, false
	}
	return insertScoreBalance(node, dir), true
}

//removeScoreBalance
func removeScoreBalance(root *Node, dir int) (*Node, bool) {
	n := root.Edges[opp(dir)]
	bal := 2*dir - 1
	switch n.Score {
	case -bal:
		root.Score = 0
		n.Score = 0
		return singleRotation(root, dir), false
	case bal:
		adjust(root, opp(dir), -bal)
		return doubleRotation(root, dir), false
	}
	root.Score = -bal
	n.Score = bal
	return singleRotation(root, dir), true
}

//removeFromRightOf
func removeFromRightOf(node *Node, data Index) (*Node, bool) {
	if node == nil {
		return nil, false
	}
	if node.Data.EqualsTo(data) {
		switch {
		case node.Edges[0] == nil:
			return node.Edges[1], false
		case node.Edges[1] == nil:
			return node.Edges[0], false
		}
		heir := node.Edges[0]
		for heir.Edges[1] != nil {
			heir = heir.Edges[1]
		}
		node.Data = heir.Data
		data = heir.Data
	}
	dir := 0
	if node.Data.LessThan(data) {
		dir = 1
	}
	var done bool
	node.Edges[dir], done = removeFromRightOf(node.Edges[dir], data)
	if done {
		return node, true
	}
	node.Score += 1 - 2*dir
	switch node.Score {
	case 1, -1:
		return node, true
	case 0:
		return node, false
	}
	return removeScoreBalance(node, dir)
}
