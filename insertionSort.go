package main

import (
	"fmt"
)

func swap(a []int, i, j int) {
	tmp := a[j]
	a[j] = a[i]
	a[i] = tmp
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j - 1]; j-- {
			swap(array, j, j - 1)
		}
	}
}

func main() {
  
  array := []int{1, 2, 3, 4, 5}
  array2 := []int{5, 4, 3, 2, 1}
  insertionSort(array)
  fmt.Println("Sorted array: ", array)
  
  insertionSort(array2)
  fmt.Println("Sorted array: ", array2)
}