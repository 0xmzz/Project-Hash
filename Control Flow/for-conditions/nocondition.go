package main

import "fmt"

func main() {
	i := 0 // iinitializer
	for {  // condition // no while in GO, you only say for
		fmt.Println(i)
		i++ // post
	}

}
