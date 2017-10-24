package heapsort

/*import "os"
import "strconv"
import "fmt"*/

//Heapsort for integers in Golang
//based on Professor Mark Allen Weiss' solution in C
//http://users.cs.fiu.edu/~weiss/

func Swap(lhs *int, rhs *int) {
	tmp := *lhs
	*lhs = *rhs
	*rhs = tmp
}

func LeftChild(in int) (out int) {
	out = 2*in + 1
	return
}

func PercDown(A []int, i int, N int) {
	var Child int
	var Tmp int

	for Tmp = A[i]; LeftChild(i) < N; i = Child {
		Child = LeftChild(i)
		if Child != N-1 && A[Child+1] > A[Child] {
			Child++
		}
		if Tmp < A[Child] {
			A[i] = A[Child]
		} else {
			break
		}
	}
	A[i] = Tmp
}

func Heapsort(A []int, N int) ([]int, int) {
	i := N
	for i = N / 2; i >= 0; i-- { //Build Heap
		PercDown(A, i, N)
	}

	for i = N - 1; i > 0; i-- {
		Swap(&A[0], &A[i]) //DeleteMax
		PercDown(A, 0, i)
	}
	return A, 1
}
