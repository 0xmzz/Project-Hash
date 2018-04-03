package main

import "fmt"

func main ()  {
    
  m := make([] string, 1, 25)
  fmt.Println(m) // [ ]
  changeMe (m)
  fmt.Println(m) // [Zubair]
  
}

func changeMe (z []string)  {  // receives a slice gets that value and assignes it to z
  z [0] = "Zubair"
  fmt.Println(z) // Zubair
}

// a slice is a reference type and the reason you can do this is because a slice is a reference itsslef meaning the address is being passed automatically 