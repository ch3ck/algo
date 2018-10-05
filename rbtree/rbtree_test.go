package rbtree

import (
	"testing"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func TestPreorder(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
	}
	tree.Preorder()
}

func TestFind(t *testing.T) {

	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	n := tree.FindIt(key(4))
	if n.Value != "dfa3" {
		t.Error("Error value")
	}
	n.Value = "bdsf"
	if n.Value != "bdsf" {
		t.Error("Error value modify")
	}
	value := tree.Find(key(5)).(string)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
	}
}
func TestIterator(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")

	it := tree.Iterator()

	for it != nil {
		it = it.Next()
	}

}

func TestDelete(t *testing.T) {
	tree := NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	for i := 1; i <= 6; i++ {
		tree.Delete(key(i))
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(key(1), "bcd4")
	tree.Clear()
	tree.Preorder()
	if tree.Find(key(1)) != nil {
		t.Error("Can't clear")
	}
}

func TestDelete2(t *testing.T) {
	tree := NewTree()
	tree.Insert(key(4), "1qa")
	tree.Insert(key(2), "2ws")
	tree.Insert(key(3), "3ed")
	tree.Insert(key(1), "4rf")
	tree.Insert(key(8), "5tg")
	tree.Insert(key(5), "6yh")
	tree.Insert(key(7), "7uj")
	tree.Insert(key(9), "8ik")
	tree.Delete(key(1))
	tree.Delete(key(2))
}
