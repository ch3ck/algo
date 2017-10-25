package list_test

import (
	"container/list"
	"fmt"
	"testing"
)

//create three lists globally
var l = list.New()
var a = list.New()
var b = list.New()

//method to print list head to tail
func printHeadToTail(l *list.List) {
	fmt.Printf("list: ")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

//method to print list tail to head
func printTailToHead(l *list.List) {
	fmt.Printf("list (reversed): ")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

func TestLinkedListMethods(t *testing.T) {

	//fmt.Println(reflect.TypeOf(l))

	//list l = [3,4,5]
	e3 := l.PushBack(3)
	l.PushBack(4)
	e5 := l.PushBack(5)

	fmt.Println("List l:")
	printHeadToTail(l)

	fmt.Println("\nInsertBefore(2, element3):")
	//pass 2, and 6 to list using
	//InsertBefore and InsertAfter methods
	l.InsertBefore(2, e3) //[(2),3,4,5]
	printHeadToTail(l)

	fmt.Println("\nInsertAfter(6, element5):")
	l.InsertAfter(6, e5) //[(2),3,4,5,(6)]
	printHeadToTail(l)

	//list a = [0,1]
	a.PushBack(0)
	a.PushBack(1)

	fmt.Println("\nList a:")
	printHeadToTail(a)

	//list b = [7,8]
	b.PushBack(7)
	b.PushBack(8)

	fmt.Println("\nList b:")
	printHeadToTail(b)

	fmt.Println("\nPushFrontList(a):")
	//append list a to the head (front) of the list
	//using PushFrontList
	l.PushFrontList(a) //([0,1],2,3,4,5)
	printHeadToTail(l)

	fmt.Println("\nPushBackList(b):")
	//append list b to the tail (back) of the list
	//using PushBackList
	l.PushBackList(b) //([0,1],2,3,4,5,6,[7,8])
	printHeadToTail(l)

	//print list length (0,1,2,3,4,5,6,7,8) => 9 elements
	fmt.Println("\nLinked list length:", l.Len())

	fmt.Println("\nPrint list in reverse:")
	//print out elements (tail to head)
	printTailToHead(l)

	eMinus1 := l.PushBack(-1)
	e9 := l.PushFront(9)

	fmt.Println("\nPushBack(-1) and PushFront(9):")
	printHeadToTail(l)

	l.MoveToFront(eMinus1)
	l.MoveToBack(e9)

	fmt.Println("\nMoveToFront(-1) and MoveToBack(9):")
	printHeadToTail(l)

	fmt.Println("\nRemove(Front()):")
	e := l.Front() //-1
	l.Remove(e)
	printHeadToTail(l)

	fmt.Println("\nMoveBefore(Front(), Back())")
	e = l.Front() //0
	f := l.Back()
	l.MoveBefore(e, f)
	printHeadToTail(l)

	fmt.Println("\nMoveAfter(Front(), Back())")
	l.MoveAfter(e, f)
	printHeadToTail(l)
}
