package btree

import (
	"encoding/json"
	"fmt"
	"testing"
)

type integerIndex int

func (ii integerIndex) LessThan(ii2 Index) bool {
	return ii < ii2.(integerIndex)
}
func (ii integerIndex) EqualsTo(ii2 Index) bool {
	return ii == ii2.(integerIndex)
}

//TestBtree
func TestBtree(t *testing.T) {
	var tree *Node
	fmt.Println("Empty Tree:")
	balancedTreeJson, _ := json.MarshalIndent(tree, "", "--")

	fmt.Println("\nInsert Tree:")
	Insert(&tree, integerIndex(4))
	if tree.Score != 4 {
		t.Fatalf("Score should be 4, got %v", tree.Score)
	}
	Insert(&tree, integerIndex(3))
	Insert(&tree, integerIndex(8))
	if tree.Score != 4 {
		t.Fatalf("Score should be 4, got %v", tree.Score)
	}
	Insert(&tree, integerIndex(6))
	Insert(&tree, integerIndex(7))
	Insert(&tree, integerIndex(9))
	if tree.Score != 6 {
		t.Fatalf("Score should be 6, got %v", tree.Score)
	}
	balancedTreeJson, _ = json.MarshalIndent(tree, "", "--")
	fmt.Println(string(balancedTreeJson))

	fmt.Println("\nRemove Tree:")
	Remove(&tree, integerIndex(4))
	Remove(&tree, integerIndex(6))
	balancedTreeJson, _ = json.MarshalIndent(tree, "", "--")
	fmt.Println(string(balancedTreeJson))
}
