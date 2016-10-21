package bst

import "testing"

func TestDelete(t *testing.T) {
	var testNode *Node
	testNode = Insert(testNode, 12)
	Insert(testNode, 1)
	Insert(testNode, 3)
	Insert(testNode, 25)

	var expectedNode *Node
	expectedNode = Insert(expectedNode, 12)
	Insert(expectedNode, 1)
	Insert(expectedNode, 3)
	Delete(testNode, 25)

	if testNode == nil {
		t.Error("Binary tree cannot be empty for deletion")
	}

	if IdenticalTrees(testNode, expectedNode) == false {
		t.Error("Deletion has not taken place")
	}

}
func TestSearch(t *testing.T) {
	var testNode *Node
	testNode = Insert(testNode, 12)
	Insert(testNode, 1)
	Insert(testNode, 3)
	Insert(testNode, 25)
	if Search(testNode, 3) == nil {
		t.Error("This is incorrect since the number is present in BST")
	}
}

func TestInsert(t *testing.T) {
	var testNode *Node
	testNode = Insert(testNode, 12)
	Insert(testNode, 1)
	Insert(testNode, 3)
	Insert(testNode, 25)

	if testNode == nil {
		t.Error("The elements haven't been inserted into the bst")
	}
}
