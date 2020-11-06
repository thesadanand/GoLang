package main

import "fmt"

func firstClasssFunctionEg() {
	fmt.Println("1st class functions eg.")

	//anonymous function to add
	addFunc := func(a, b int) int {
		fmt.Println("addition of ", a, b)
		return a + b
	}
	//anonymous function to sub
	subFunc := func(a, b int) int {
		fmt.Println("subtraction of ", a, b)
		return a - b
	}

	functionsArr := [2]func(int, int) int{addFunc, subFunc}

	for _, function := range functionsArr {
		fmt.Println(function(20, 10))
	}
}

func getSqValue() func(a int) int {
	return func(a int) int {
		return a * a
	}
}

func calculate(a, b int, functionToApply func(c, d int) int) int {
	return functionToApply(a, b)
}

func higherOderFunctionEg() {

	mutliplication := func(a, b int) int {
		return a * b
	}

	division := func(a, b int) int {
		return a / b
	}

	//calculate function is taking a function as an arg
	fmt.Println("product of 10 and 20", calculate(10, 20, mutliplication))
	fmt.Println("division of 10 and 20", calculate(10, 20, division))

	//returning a function in go
	sqFunc := getSqValue()
	fmt.Println("square of 10 is : ", sqFunc(10))
}

func main() {
	firstClasssFunctionEg()
	higherOderFunctionEg()
}
