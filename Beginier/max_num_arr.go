package main 

import (
 "fmt"
 "math/rand"
  "time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
  var arr[99] int
  var max int
  var n int 
  fmt.Println("Enter n: ")
  fmt.Scanln(&n)
  for i:=0;i<n;i++ {
	 arr[i]= rand.Intn(99)
  }
  for i:=0;i<n;i++ {
  	fmt.Print(" Arr is:  ",arr[i]," ")
  }
  for i:=0;i<n;i++ {
	 if(max < arr[i]){
	    max = arr[i]
	 }
  }
  fmt.Printf(" Max number is:  ",max)
}

