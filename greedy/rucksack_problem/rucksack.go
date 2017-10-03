/*Package qsort -- Greedy problem: Rucksack problem implementation: */
/* This Algorithm  gives you the best result for */
/* filling a rucksack, with book with a value. */
/* It gets the max value of the combination of books*/
/* that could fit in the bag. */
/* Author: Carlos Duran (@duxy1996)*/

package main

import "fmt"
import "math"

/*Create struct to simulate a tupla*/
type Pair struct {
    a, b float64
}

func rucksack(books []Pair, weight float64)(result float64){
    var yes float64
    var no float64
    var book Pair
    var books2 []Pair
    if (len(books) == 0) || (weight == 0) {
        result = 0
        return 
    } else{
        book = books[0]       
        for i := 1; i < len(books); i++ {
            books2 = append(books2, books[i])
        }               
        if book.b <= weight {
            yes = book.a + rucksack(books2,weight - book.b)
        } else {
            yes = 0
        }        
        no = 0 + rucksack(books2, weight)
        result = math.Max(yes,no)
        return
    }

}

func main() {
     var result float64
     test := []Pair{Pair{40, 16},Pair{41, 16},Pair{42, 30}}
     result = rucksack(test, 30)
     fmt.Println("Result", result)

}