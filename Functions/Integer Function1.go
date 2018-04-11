package main

import "fmt"

func half (n int) (float64, bool) {			//  the first ( n is a variable that is an integer)
		return float64(n) / 2, n%2 == 0 // % is the remainder, its an expression that will evaluate to a type bool

}

func main () {
	 	x, n := half(5)
	 fmt.Println(x, n)

}