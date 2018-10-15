package bubblesort
//
//import (
//	"math/rand"
//	"reflect"
//	"sort"
//	"testing"
//	"time"
//)
//
//// TestBubbleSort inspired in mergesort_test
//func TestBubbleSort(t *testing.T) {
//	empty := []int{}
//	resultEmpty := bubbleSort(empty)
//
//	if !reflect.DeepEqual(empty, resultEmpty) {
//		t.Error("ERROR!", "Expected:", empty, "Result:", resultEmpty)
//		return
//	}
//
//	unitary := []int{42}
//	resultUnitary := bubbleSort(unitary)
//
//	if !reflect.DeepEqual(unitary, resultUnitary) {
//		t.Error("ERROR!", "Expected:", unitary, "Result:", resultUnitary)
//		return
//	}
//
//	rand.Seed(int64(time.Now().Nanosecond()))
//	for test := 0; test < 10; test++ {
//		array := make([]int, 100)
//		for i := 0; i < 100; i++ {
//			array[i] = rand.Int()
//		}
//
//		expected := make([]int, 100)
//		copy(expected, array)
//		sort.Ints(expected)
//
//		result := bubbleSort(array)
//
//		if !reflect.DeepEqual(result, expected) {
//			t.Error("ERROR!", "Expected:", expected, "Result:", result)
//			return
//		}
//	}
//}
