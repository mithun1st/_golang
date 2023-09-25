package main

import (
	"fmt"
)

func main() {
	var arr [5]int = [5]int{44, 55, 66, 77, 88}

	var newSlice1 = make([]int, 0)
	newSlice1 = arr[0:3]

	var newSlice2 []int = []int{}
	newSlice2 = arr[3:]

	fmt.Println(arr)
	fmt.Println(newSlice1)
	fmt.Println(newSlice2)

	newSlice2 = append(newSlice2, newSlice1...)
	fmt.Println(newSlice2)
}
