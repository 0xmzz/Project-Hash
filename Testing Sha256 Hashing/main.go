package main

import "fmt"

func main() {
  a := 44
  fmt.Println("a", a)
  fmt.Println("a's memory address", &a)
  fmt.Printf("%d \n", &a)
}