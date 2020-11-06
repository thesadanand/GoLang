package main

import (
	"fmt"
	"strings"
)

func ceaserCipher(input string, by int) string {

	output := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] + by
		//fmt.Println(input[i])
	}
	fmt.Println(output)
	return strings.Join(output, " ")
}

func main() {
	input := "madam"
	var shiftBy int = 3
	fmt.Println("input", input)
	output := ceaserCipher(input, shiftBy)
	fmt.Println("ceaser Cipher", output)

}
