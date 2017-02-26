package ludecomposition

import (
	"math"
	"reflect"
	"testing"
)

// Holds input and output for a testcase
type testdata struct {
	matrix [][]float64
	l      [][]float64
	u      [][]float64
	perm   []int
}

var tests = []testdata{
	{[][]float64{{4.0, 3.0}, {6.0, 3.0}},
		[][]float64{{1.0, 0.0}, {2.0 / 3.0, 1.0}},
		[][]float64{{6.0, 3.0}, {0.0, 1.0}},
		[]int{1, 0}},
	{[][]float64{{3.0, 6.0}, {-3.0, -5.0}},
		[][]float64{{1.0, 0.0}, {-1.0, 1.0}},
		[][]float64{{3.0, 6.0}, {0.0, 1.0}},
		[]int{0, 1}},
	{[][]float64{{1.0, 2.0, 1.0}, {3.0, 3.0, 1.0}, {0.0, 9.0, 1.0}},
		[][]float64{{1.0, 0.0, 0.0}, {0.0, 1.0, 0.0}, {3.0 / 9.0, 1.0 / 9.0, 1.0}},
		[][]float64{{3.0, 3.0, 1.0}, {0.0, 9.0, 1.0}, {0.0, 0.0, 5.0 / 9.0}},
		[]int{1, 2, 0}},
}

// Compares two matrices for equality, considering floating point values
// with absolute difference less than 1e-9 as equal.
func areEqual(a, b [][]float64) bool {
	var EPSILON float64 = 1e-9

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if math.Abs(a[i][j]-b[i][j]) > EPSILON {
				return false
			}
		}
	}

	return true
}

// Tests the algorithm for LU decomposition. Note that, assuming L has only ones on
// its main diagonal, the solution is unique.
func TestLuDecomposition(t *testing.T) {
	for _, test := range tests {
		l, u, perm, err := FindLuDecomposition(test.matrix)

		if err != nil {
			t.Error("Error occurred:", err)
		}

		if !areEqual(l, test.l) {
			t.Error(
				"For input", test.matrix,
				"expected L", test.l,
				"got", l,
			)
		}

		if !areEqual(u, test.u) {
			t.Error(
				"For input", test.matrix,
				"expected U", test.u,
				"got", u,
			)
		}

		if !reflect.DeepEqual(perm, test.perm) {
			t.Error(
				"For input", test.matrix,
				"expected permutation", test.perm,
				"got", perm,
			)
		}
	}
}
