package main

import "fmt"

func main() {
	switch "zubair" {

	case "daniel":
		fmt.Println("Hey daniel")
	case "zubair":
		fmt.Println("Hey Zubair")
		fallthrough
	case "jaggy":
		fmt.Println("Hey, Jaggers")
	default:
		fmt.Println("wasa wasa wasaaaaaaap")
	}

}
