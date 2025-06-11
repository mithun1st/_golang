package main

import (
	"fmt"
)

func main() {

	//* Array
	var nums1 [5]int
	nums1[0] = 1
	fmt.Println(nums1)
	fmt.Println("Len:", len(nums1), "Cap:", cap(nums1))

	var nums2 [3]int = [...]int{1, 2, 3}
	fmt.Println(nums2)
	fmt.Println("Len:", len(nums2), "Cap:", cap(nums2))

	var nums3 []int = []int{11, 0, 0, 0, 0}
	nums3[1] = 22
	fmt.Println(nums3)
	fmt.Println("Len:", len(nums3), "Cap:", cap(nums3))
	for index, value := range nums3 {
		fmt.Printf("Index:%d Value:%v\n", index, value)
	}

	var nums4 = [5]int{1, 2, 3}
	fmt.Println(nums4)
	fmt.Println("Len:", len(nums4), "Cap:", cap(nums4))

	var strArr []rune = []rune("Mithun")
	fmt.Println(strArr)
	fmt.Println("Len:", len(strArr), "Cap:", cap(strArr))

	//* Two dimension array
	var twoDimentionArray [3][4]int = [3][4]int{
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}

	for i, _ := range twoDimentionArray {
		for j, _ := range twoDimentionArray[i] {
			fmt.Printf("*%d*", twoDimentionArray[i][j])
		}
		fmt.Println()
	}

	//* Map
	var map1 map[int]string
	fmt.Println(map1)

	var map2 map[int]string = map[int]string{101: "Apple", 102: "Ball"}
	fmt.Println(map2)

	fmt.Println(map2[101])

	value1, hasExist1 := map2[100]
	if hasExist1 {
		fmt.Println(value1)
	} else {
		fmt.Println("Not Found")
	}

	value2, hasExist2 := map2[102]
	if hasExist2 {
		fmt.Println(value2)
	} else {
		fmt.Println("Not Found")
	}

	// map add and delete
	map2[103] = "Cat"
	fmt.Println(map2)
	delete(map2, 103)
	fmt.Println(map2)
	clear(map2)
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Printf("Key:%v Value:%v\n", key, value)
	}

	/*
		[1 0 0 0 0]
		Len: 5 Cap: 5
		[1 2 3]
		Len: 3 Cap: 3
		[11 22 0 0 0]
		Len: 5 Cap: 5
		Index:0 Value:11
		Index:1 Value:22
		Index:2 Value:0
		Index:3 Value:0
		Index:4 Value:0
		map[]
		map[101:Apple 102:Ball]
		Apple
		Not Found
		Ball
		map[101:Apple 102:Ball 103:Cat]
		map[101:Apple 102:Ball]
		map[]
	*/
}
