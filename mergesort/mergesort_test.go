package mergesort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	empty := []int{}
	resultEmpty := MergeSort(empty)

	if !reflect.DeepEqual(empty, resultEmpty) {
		t.Error("ERROR!", "Expected:", empty, "Result:", resultEmpty)
		return
	}

	unitary := []int{1921}
	resultUnitary := MergeSort(unitary)

	if !reflect.DeepEqual(unitary, resultUnitary) {
		t.Error("ERROR!", "Expected:", unitary, "Result:", resultUnitary)
		return
	}

	rand.Seed(int64(time.Now().Nanosecond()))
	for test := 0; test < 10; test++ {
		array := make([]int, 100)
		for i := 0; i < 100; i++ {
			array[i] = rand.Int()
		}

		expected := make([]int, 100)
		copy(expected, array)
		sort.Ints(expected)

		result := MergeSort(array)

		if !reflect.DeepEqual(result, expected) {
			t.Error("ERROR!", "Expected:", expected, "Result:", result)
			return
		}
	}
}
