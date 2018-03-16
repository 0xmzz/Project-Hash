package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
		fmt.Println("end of innerloop")
	}
	fmt.Println("end of outerloop")
}
