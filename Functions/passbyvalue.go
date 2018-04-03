package main

import "fmt"

func main ()  {
    
  age := 44
  fmt.Println(&age) // Will give you the momeory address of where age is stored
  
  changeMe(&age)
  
 // fmt.Println(&age) // add of where age is stored
  fmt.Println(age) // 24
}

func changeMe (z *int)  {   // * is part of a type, a pointer to an int
  
  fmt.Println(z) //address where age is stored
  fmt.Println(*z)//44, also * is an operater
  *z = 24
  
  fmt.Println(z)
  //fmt.Println(*z)
  
  
}