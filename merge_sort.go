package main

import "fmt"

//merge function that takes two integer arrays and join them to return a sorted array
func merge(left, right []int) []int {

  merged_arr := make([]int, 0)

  for len(left) > 0 || len(right) > 0 {
    if len(left) > 0 && len(right) > 0 {
      if left[0] <= right[0] {
        merged_arr = append(merged_arr, left[0])
        left = left[1:len(left)]
      } else {
        merged_arr = append(merged_arr, right[0])
        right = right[1:len(right)]
      }
    } else if len(left) > 0 {
      merged_arr = append(merged_arr, left[0])
      left = left[1:len(left)]
    } else if len(right) > 0 {
      merged_arr = append(merged_arr, right[0])
      right = right[1:len(right)]
    }
  }

  return merged_arr
}

//merge sort algorithm using above merge function
func mergeSort(arr []int) []int {

  if len(arr) <= 1 {
    return arr
  }
  
  left := make([]int, 0)
  right := make([]int, 0)
  mid := len(arr) / 2

  for i, x := range arr {
    switch {
    case i < mid:
      left = append(left, x)
    case i >= mid:
      right = append(right, x)
    }
  }

  left = mergeSort(left)
  right = mergeSort(right)

  return merge(left, right)
}
func main() {
  
  arr := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
  fmt.Println("Original Array: ", arr)
  arr = mergeSort(arr)
  fmt.Println("Sorted Array: ", arr)
}
