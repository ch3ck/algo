// Package provides an algorithm for constructing the LU decomposition
// (with pivoting) of a given matrix.
package ludecomposition

import (
	"errors"
	"math"
)

// Applies pivoting to the input. If the given matrix is invertible,
// after pivotization there are no zeros on the main diagonal.
// Returns the permutation that was applied on the rows of the matrix.
func pivotize(matrix [][]float64) (perm []int) {
	n := len(matrix)
	perm = make([]int, n)

	// Initialize perm to identity permutation
	for i := range perm {
		perm[i] = i
	}

	for i := range matrix {
		max := matrix[i][i]
		pos := i

		// Extract maximum from the values in matrix[i:n-1][i]
		for j := i; j < n; j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
				pos = j
			}
		}

		// Swap rows bringing the maximum value to the main diagonal
		if pos != i {
			perm[i], perm[pos] = perm[pos], perm[i]
			matrix[i], matrix[pos] = matrix[pos], matrix[i]
		}
	}

	return perm
}

// Returns an n x n identity matrix.
func identityMatrix(n int) (matrix [][]float64) {
	matrix = make([][]float64, n)

	for i := range matrix {
		matrix[i] = make([]float64, n)

		for j := range matrix[i] {
			matrix[i][j] = 0.0
		}

		matrix[i][i] = 1.0
	}

	return
}

// Implementation of the LU decomposition algorithm.
// Returns L, U and perm, where L and U form the decomposition,
// and perm is the row permutation applied during pivoting.
// The L matrix has only ones on it's main diagonal, which makes the returned solution
// uniquely defined.
// Raises an error if the input matrix is non-invertible.
func FindLuDecomposition(matrix [][]float64) ([][]float64, [][]float64, []int, error) {
	var EPSILON float64 = 1e-9

	n := len(matrix)

	l := identityMatrix(n)
	u := make([][]float64, n)

	// Initialize u as a deep copy of the input matrix
	for i := range u {
		u[i] = make([]float64, n)
		copy(u[i], matrix[i])
	}

	perm := pivotize(u)

	for i := range u {
		// If after pivotization there are values equal to zero on the main diagonal
		// it means that the input matrix is singular, and thus there is no solution
		if math.Abs(u[i][i]) < EPSILON {
			return nil, nil, nil, errors.New("matrix is singular")
		}
	}

	for i := range u {
		for j := i + 1; j < n; j++ {
			ratio := u[j][i] / u[i][i]

			// Zero out u[j][i] by subtracting from the j'th row the i'th row
			// multiplied by a suitable coefficient
			for k := 0; k < n; k++ {
				u[j][k] -= ratio * u[i][k]
			}

			// Permorm an analogous column-wise operation
			// in order to keep the L*U product constant
			for k := 0; k < n; k++ {
				l[k][i] += ratio * l[k][j]
			}
		}
	}

	return l, u, perm, nil
}
