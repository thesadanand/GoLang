package main

import (
	"fmt"
)

var strarr = []string{"abc", "pqr", "zyx", "vbn"}
var intarr = []int{1, 2, 3, 4}
var mapone = map[int]string{}
var maptwo = map[string]interface{}{}

func main() {
	for i := 0; i != 4; i++ {
		fmt.Println(intarr[i], "\t", strarr[i])
		mapone[intarr[i]] = strarr[i]
		maptwo[strarr[i]] = mapone
	}
	fmt.Println(mapone)
	fmt.Println(maptwo)
}
