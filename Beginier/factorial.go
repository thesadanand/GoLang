package main 

import (
	"fmt"
)

func main() {
	var n int = 5
    var fact int = 1
	for i:=1;i<=n;i++{
		fact = fact * i
	}
	fmt.Print(" factorial is:  ",fact)

}

