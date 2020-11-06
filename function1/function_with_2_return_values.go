package main

import "fmt"

func rectangle(l int, b int) (area int, perimeter int) {
	perimeter = 2 * (l + b)
	area = l * b
	return
}

func main() {
	var a, p int
	a, p = rectangle(12, 13)
	fmt.Println("perimeter: ", p)
	fmt.Println("area: ", a)
}
