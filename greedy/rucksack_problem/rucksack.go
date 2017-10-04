/*Package qsort -- Greedy problem: Rucksack problem implementation: */
/* This Algorithm  gives you the best result for */
/* filling a rucksack, with book with a value. */
/* It gets the max value of the combination of books*/
/* that could fit in the bag. */
/* Author: Carlos Duran (@duxy1996)*/

/*Algorithm explication*/
/*it makes a recursive calling with*/
/*two its two gays, put the book or not put the book*/
/*it happens until there is no book o dont fit more*/
/*then in each call return the best value*/

package rucksack

import "math"

/*Create struct to simulate a tupla*/
/*money is the first element (Value)*/
/*weight is the first element (Weight)*/
type Pair struct {
	money, weight float64
}

/*This algorithm return the max value possible*/
/*with the combination of books that could fit in the rucksack*/
/*The imputs are a list of books and the max weight (float64)*/
/*The result is a float64*/
func rucksack(books []Pair, weight float64) (result float64) {
	/*define the variable yes: contains the solution with the book*/
	/*define the variable no: contains the solution without the book*/
	/*define the variable book: to save the first book of the list*/
	/*define the variable books_copy: a copy of the list of books*/
	var yes float64
	var no float64
	var book Pair
	var books_copy []Pair
	/*Base case: if there is no book or dont fit more, return 0*/
	if (len(books) == 0) || (weight == 0) {
		result = 0
		return
	} else {
		/*get the first book*/
		book = books[0]
		/*remove the first book*/
		for i := 1; i < len(books); i++ {
			books_copy = append(books_copy, books[i])
		}
		/*if the book fits, try to fit another*/
		if book.weight <= weight {
			yes = book.money + rucksack(books_copy, weight-book.weight)
		} else {
			yes = 0
		}
		/*fit other book*/
		no = 0 + rucksack(books_copy, weight)
		/*get the max value from both recursive calls and return the max*/
		result = math.Max(yes, no)
		return
	}

}
