/* Testing the Qsort() function in the qsort package */
/* Author: Alangi Derick (@d3r1ck), Date: 10/10/2016 */

package qsort

import (
	"reflect"
	"testing"
)

/* Test pair for Qsort() algorithm */
type testpair struct {
	unsortedSlice []int
	sortedSlice   []int
}

/* Test instance of testpair the unsorted and sorted slices */
var tests = []testpair{
	{
		unsortedSlice: []int{3, 2, 1, 5, 6, 3, 7, 8, 9},
		sortedSlice:   []int{1, 2, 3, 3, 5, 6, 7, 8, 9},
	},
	{
		unsortedSlice: []int{1, 4, 2, 3},
		sortedSlice:   []int{1, 2, 3, 4}},
}

/* Test case to test the Qsort function */
func TestQsort(t *testing.T) {
	for _, pair := range tests {
		sorted := Qsort(pair.unsortedSlice)

		if !reflect.DeepEqual(sorted, pair.sortedSlice) {
			t.Error(
				"For", pair.unsortedSlice,
				"expected", pair.sortedSlice,
				"got", sorted,
			)
		}
	}
}
