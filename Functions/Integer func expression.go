package main

import "fmt"



func main () {

	half := func (n int) (int, bool) {
		return n / 2, n%2 == 0

}
	 	//x, even := half(4)
	 fmt.Println(half(100))

}


