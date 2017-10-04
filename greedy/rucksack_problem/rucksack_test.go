/* Testing the rucksack() function in the rucksack package */
/* Author: Carlos Duran (@duxy1996), Date: 03/10/2017 */

package rucksack

import (
    "reflect"
    "testing"
)


/* Test pair for rucksack() algorithm */
type testpair struct {
    books []Pair
    max_value float64
}

var tests = []testpair{
    {[]Pair{Pair{40, 16},Pair{41, 16},Pair{42, 30}},42},
    {[]Pair{Pair{40, 16},Pair{41, 16},Pair{42, 30}},42},
}

/* Test case to test the Qsort function */
func Testrucksack(t *testing.T) {
    for _, pair := range tests {
        value := rucksack(pair.books,30)

        if !reflect.DeepEqual(value, pair.max_value) {
            t.Error(
                "For", pair.books,
                "expected", pair.max_value,
                "got", value,
            )
        }
    }
}