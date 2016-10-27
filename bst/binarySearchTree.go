package bst

import (
	"fmt" //importing required packages
)

type Node struct { //Node represents node of a Binary Search Tree
	value int
	left  *Node
	right *Node
}

func Insert(root *Node, key int) *Node { //The following function inserts elements into Binary Search Tree
	if root == nil {
		root = &Node{key, nil, nil} //root node is created with left and right child as nil
	} else if key < root.value {
		root.left = Insert(root.left, key) //Using recursion to maintain the tree as BST
	} else if key > root.value {
		root.right = Insert(root.right, key)
	}
	return root
}

func Search(root *Node, key int) *Node { //The following function searches a given key in a BST
	if root == nil {
		return root
	}

	if root.value == key { //If the key is not found a nil pointer is returned
		return root
	}
	if root.value < key {
		return Search(root.right, key)
	}
	return Search(root.left, key)

}
func Minval(root *Node) *Node { //The following function is used in Delete function to find the inorder successor
	var current *Node
	current = root
	for current.left != nil {
		current = current.left
	}
	return current
}
func Delete(root *Node, key int) *Node { //The following function deletes the key in BST and returns the new root
	if root == nil {
		//base case
		return root
	}
	if key < root.value { //If the key to be deleted is smaller than the root's key,the key lies in the left subtree
		root.left = Delete(root.left, key)
	} else if key > root.value { //If the key to be deleted is greater than the root's key,the key lies in the right subtree
		root.right = Delete(root.right, key)
	} else { //node with only one child
		if root.left == nil {
			var temp *Node = root.right
			return temp
		} else if root.right == nil {
			var temp *Node = root.left
			return temp
		} //node with two children
		var temp *Node
		temp = Minval(root.right)                   //Get inorder successor
		root.value = temp.value                     //Get the inorder successor's value
		root.right = Delete(root.right, temp.value) //Delete the inorder successor
	}
	return root
}
func PreorderTraverse(root *Node) { //The following function does preorder traversal of BST,i.e,Root,left,Right nodes.
	if root == nil {
		return
	}
	fmt.Println(root.value)
	PreorderTraverse(root.left)
	PreorderTraverse(root.right)
}

func InorderTraverse(root *Node) { //The following function does inorder traversal of BST,i.e,Left,Root,Right nodes.
	if root == nil {
		return
	}
	InorderTraverse(root.left)
	fmt.Println(root.value)
	InorderTraverse(root.right)
}

func PostorderTraverse(root *Node) { //The following function does postorder traversal of BST,i.e,Left,Right,Root nodes.
	if root == nil {
		return
	}
	PostorderTraverse(root.left)
	PostorderTraverse(root.right)
	fmt.Println(root.value)
}
func IdenticalTrees(a *Node, b *Node) bool { //The following function checks whether two trees are indentical or not
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		return (a.value == b.value && IdenticalTrees(a.left, b.left) && IdenticalTrees(a.right, b.right))
	}
	return false
}
