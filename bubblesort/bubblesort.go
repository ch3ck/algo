// Package bubblesort provides implementation for bubble sort algorithm
package bubblesort

// Swap function changes two integers
func Swap(lhs *int, rhs *int) {
	tmp := *lhs
	*lhs = *rhs
	*rhs = tmp
}

// Bubblesort function receives an array and return the sorted array
// O(N^2)
func bubbleSort(A []int) []int {
	N := len(A)
	result := make([]int, len(A))
	i := 0
	j := 0
	for i = 0; i < N; i++ {
		result[i] = A[i]
	}
  
	for i = 0; i < N-1; i++ {
		for j = 0; j < N-1; j++ {
			if result[j] > result[j+1] {
				Swap(&result[j], &result[j+1])
			}
		}
	}
	return result
}
