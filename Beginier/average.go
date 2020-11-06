package main 

import (
  "fmt"
)

func main() {
  var arr[99] int
  var avg,sum,tmp int
  fmt.Print("Enter number of elements: ")
  fmt.Scanln(&tmp)
  for i:=0;i<tmp;i++ {
	 fmt.Print("Enter number: ",i+1," ")
	 fmt.Scanln(&arr[i])
	 sum += arr[i]
  }
  avg = sum/tmp
  fmt.Print("Average is: ",avg)
}

