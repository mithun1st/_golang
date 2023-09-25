package main

import (
	"fmt"
)

func main() {
	// 2d array
	var mattrix [4][3]string = [4][3]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"X", "Y", "Z"},
	}

	for _, ary := range mattrix {
		for _, e := range ary {
			fmt.Print(e, " ")
		}
		fmt.Println()
	}
	fmt.Println(mattrix)
	fmt.Println()

	// 2d slice
	var mattrixSlice [][]bool = [][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, true},
	}

	for _, ary := range mattrixSlice {
		for _, e := range ary {
			fmt.Print(e, " ")
		}
		fmt.Println()
	}
	fmt.Println(mattrixSlice)
}
