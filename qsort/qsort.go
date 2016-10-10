/* Quick Sort Algorithm Data structure implemented */
/* This Algorithm takes an slice of integers and   */
/* sort using the qsort() algorithm below 		   */
/* Author: Alangi Derick (@d3r1ck) 				   */

/* Create the package */
package qsort

import "math/rand"

/* As GoLang standard, the function in a package should */
/* start with Upper case letter like; Qsort()		    */
func Qsort(slice []int) []int {

	/* Check if the slice has just a single cell */
	/* If it does have just a single cell, return */
	/* the slice. 								 */
	if len(slice) < 2 {
		return slice
	}

	left, right := 0, len(slice)-1

	/* Pick a pivot in the slice */
	pivotIndex := rand.Int() % len(slice)

	/* Move pivot to the right of slice */
	slice[pivotIndex], slice[right] = slice[right], slice[pivotIndex]

	/* Move elements smaller than pivot to the left */
	for i := range slice {
		if slice[i] < slice[right] {
			slice[i], slice[left] = slice[left], slice[i]
			left++
		}

	}

	/* Place the pivot after the last smaller element */
	slice[left], slice[right] = slice[right], slice[left]

	/* Do a recursive call to the Qsort function */
	/* with new values to sort entire slice */
	Qsort(slice[:left])
	Qsort(slice[left+1:])

	return slice
}
