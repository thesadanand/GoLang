package main 

import (
 "fmt"
)

func main() {
  fmt.Print("Enter a number: ")
  var n int
  fmt.Scanln(&n)
  if(n%2==0){
	  fmt.Print(n," is an even number")
  }else{
	   fmt.Print(n," is an odd number")
  }
}

