package main

import (
	"fmt"
	"slices"
)

func main() {
	//* Append
	var arr1 []int = []int{11, 22, 33}
	fmt.Println(arr1)
	fmt.Println("Len:", len(arr1), "Cap:", cap(arr1))
	arr1 = append(arr1, 14)
	fmt.Println(arr1)
	fmt.Println("Len:", len(arr1), "Cap:", cap(arr1))

	//* Make
	var arr2 []int = make([]int, 5, 10)
	fmt.Println(arr2)
	fmt.Println("Len:", len(arr2), "Cap:", cap(arr2))

	arr2 = append(arr2, []int{1, 2, 3, 4, 5}...)
	fmt.Println(arr2)
	fmt.Println("Len:", len(arr2), "Cap:", cap(arr2))

	//* Slice
	var arr3 []int = make([]int, 0)
	arr3 = []int{10, 20, 30, 40, 50}
	fmt.Println(arr3)
	fmt.Println("Len:", len(arr3), "Cap:", cap(arr3))

	arr3 = append(arr3[:1], arr3[2:]...) //delete index 1
	fmt.Println(arr3)
	fmt.Println("Len:", len(arr3), "Cap:", cap(arr3))

	//* Is equal
	var arr4 []int = []int{1, 2, 3}
	var arr5 []int = []int{1, 2, 3}

	var isEqual bool = slices.Equal(arr4, arr5)
	fmt.Println(isEqual)

	/*
		[11 22 33]
		Len: 3 Cap: 3
		[11 22 33 14]
		Len: 4 Cap: 6
		[0 0 0 0 0]
		Len: 5 Cap: 10
		[0 0 0 0 0 1 2 3 4 5]
		Len: 10 Cap: 10
		[10 20 30 40 50]
		Len: 5 Cap: 5
		[10 30 40 50]
		Len: 4 Cap: 5
		true
	*/
}
