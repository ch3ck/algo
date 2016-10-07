package main

import "fmt"

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

func Heapsort(A []int, N int) {
	i := N
	for i = N / 2; i >= 0; i-- { //Build Heapsort
		PercDown(A, i, N)
	}

	for i = N - 1; i > 0; i-- {
		Swap(&A[0], &A[i]) //DeleteMax
		PercDown(A, 0, i)
	}
}

func main() {
	array := []int{3, 1, 4, 1, 5, 9, 2, 6}
	fmt.Printf("before: %v\n", array)
	Heapsort(array, len(array)) //pass heap and length count
	fmt.Printf("after: %v", array)

	/* Test for Swap func
	   j, k := 10, 5
	   pj, pk := &j, &k
	   fmt.Printf("before: J:%d, K:%d\n", *pj, *pk)
	   Swap(pj,pk)
	   fmt.Printf("after: J:%d, K:%d\n", *pj, *pk)
	*/
}
