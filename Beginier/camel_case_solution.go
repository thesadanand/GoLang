package main

import (
	"fmt"
	"strings"
)

func camelCaseSolution(s string) int32 {
	words := 1
	for _, ch := range s {
		str := string(ch)
		if strings.ToUpper(str) == str {
			words++
		}
	}
	return int32(words)
}

func main() {
	var input string = "noWomenNoCry"
	tokens := camelCaseSolution(input)
	fmt.Println(input)
	fmt.Println(tokens)
}
