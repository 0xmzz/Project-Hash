package main

import "fmt"

func main () {
  fmt.Println(greet("Zubair", "Zia"))
  }

func greet (fname, lname string) string {
  return fmt.Sprint(lname, fname)
}