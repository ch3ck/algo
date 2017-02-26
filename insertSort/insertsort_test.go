package insertsort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {

	array := []int{1, 6, 2, 4, 9, 0, 5, 3, 7, 8}
	fmt.Println("not sorted array: ", array)
	insertSort(array)
	fmt.Println("just sorted array: ", array)
}
