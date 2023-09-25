package main

import (
	"fmt"
	"reflect"
)

func sum(num []int) int {
	var sum int
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func main() {
	var myArray [8]int
	myArray = [8]int{55, 66, 77, 88, 99, 11, 33, 44}

	fmt.Println(sum(([]int{55, 66, 77, 88, 99, 11, 33, 44})[:]))

	//slice
	fmt.Println(myArray)
	mySlice := myArray[:]
	fmt.Println(mySlice)
	mySlice = myArray[1:4] //slice (66 to 88)
	fmt.Println(len(mySlice))
	fmt.Println(mySlice)
	mySlice = myArray[1:5] //slice (66 to 99)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(reflect.TypeOf(mySlice))

	fmt.Println("----------------------")

	//build slice by make
	var menualSlice = make([]int, 3, 5)
	menualSlice = []int{11, 22, 33, 44, 55, 66, 77, 88}

	fmt.Println(menualSlice)

	fmt.Println(len(menualSlice))
	fmt.Println(cap(menualSlice))
	menualSliceWithoutCap := make([]int, 5)
	fmt.Println(len(menualSliceWithoutCap))
	fmt.Println(cap(menualSliceWithoutCap))

	//
	menualSlice1 := make([]int, 1, 10)
	menualSlice1 = []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(menualSlice1)

	fmt.Println(reflect.TypeOf(menualSlice1))

	//append
	var slice []string = make([]string, 0)
	slice = []string{"a", "b", "c"}
	slice2 := []string{"x", "y"}
	array3 := [2]string{"z", "z1"}
	slice = append(slice, slice2...)
	slice = append(slice, array3[:]...)
	fmt.Println(slice)

	longSlice := make([]int, 0)
	for i := 100; i < 200; i++ {
		longSlice = append(longSlice, i)
	}
	fmt.Println(longSlice)

	//range
	for i, e := range longSlice {
		fmt.Println(i, " - ", e)
	}

	//delete 5 index
	longSlice = append(longSlice[:5], longSlice[5+1:]...)
	fmt.Println(longSlice)

	//clear
	longSlice = nil
	fmt.Println(longSlice)
}
