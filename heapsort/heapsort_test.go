package heapsort_test

import (
	"fmt"
	"testing"

	"github.com/Ch3ck/Algo/heapsort"
)

func TestHeapsort(t *testing.T) {
	var data = []int{3, 1, 4, 1, 5, 9, 2, 6} //create data
	fmt.Println("input:  ", data)
	var expected = []int{1, 1, 2, 3, 4, 5, 6, 9}
	var result, flag = heapsort.Heapsort(data, len(data))

	if flag != 1 {
		fmt.Print("Something went wrong")
		return
	}

	//check if result is null
	if result == nil {
		t.Error("returned value is nil")
		return
	}

	//result is not null, we perform checks
	if len(result) != len(expected) {
		t.Error("sorted data has different length from input data")
		return
	} else { //lengths are okay, we compare output values with expected values
		for i := range result {
			if result[i] != expected[i] {
				t.Error("failed at sorting")
			} else {
				fmt.Println("expect: ", expected)
				fmt.Println("output: ", result)
				fmt.Println("success")
				return
			}
		}
	}
}
