package insertsort

import (
	"reflect"
	"testing"
)

/* Test pair for InsertSort() algorithm */
type testpair struct {
	unsortedSlice []int
	sortedSlice   []int
}

/* Test instance of testpair the unsorted and sorted slices */
var tests = []testpair{
	{[]int{3, 2, 1, 5, 6, 3, 7, 8, 9}, []int{1, 2, 3, 3, 5, 6, 7, 8, 9}},
	{[]int{1, 4, 2, 3}, []int{1, 2, 3, 4}},
	{[]int{4, 5, 2, 7, 6, 8, 7}, []int{2, 4, 5, 6, 7, 7, 8}},
}

/* Test case to test the InsertSort function */
func TestInsertSort(t *testing.T) {
	for _, pair := range tests {
		sorted := InsertSort(pair.unsortedSlice)

		if !reflect.DeepEqual(sorted, pair.sortedSlice) {
			t.Error(
				"For", pair.unsortedSlice,
				"expected", pair.sortedSlice,
				"got", sorted,
			)
		}
	}
}
