// Provides an algorithm for constructing the LU decomposition (with pivoting) of a given matrix
package ludecomposition

import (
	"errors"
	"math"
)

func pivotize(matrix [][]float64) (perm []int) {
	n := len(matrix)
	perm = make([]int, n)

	for i := range perm {
		perm[i] = i
	}

	for i := range matrix {
		max := matrix[i][i]
		pos := i

		for j := i; j < n; j++ {
			if matrix[j][i] > max {
				max = matrix[j][i]
				pos = j
			}
		}

		if pos != i {
			perm[i], perm[pos] = perm[pos], perm[i]
			matrix[i], matrix[pos] = matrix[pos], matrix[i]
		}
	}

	return perm
}

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

func FindLuDecomposition(matrix [][]float64) ([][]float64, [][]float64, []int, error) {
	var EPSILON float64 = 1e-9

	n := len(matrix)

	l := identityMatrix(n)
	u := make([][]float64, n)

	for i := range u {
		u[i] = make([]float64, n)
		copy(u[i], matrix[i])
	}

	perm := pivotize(u)

	for i := range u {
		if math.Abs(u[i][i]) < EPSILON {
			return nil, nil, nil, errors.New("matrix is singular")
		}
	}

	for i := range u {
		for j := i + 1; j < n; j++ {
			ratio := u[j][i] / u[i][i]

			for k := 0; k < n; k++ {
				l[k][i] += ratio * l[k][j]
			}

			for k := 0; k < n; k++ {
				u[j][k] -= ratio * u[i][k]
			}
		}
	}

	return l, u, perm, nil
}
