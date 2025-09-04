package main

import (
	"fmt"
)

func setZeroWithoutPointer(value int) {
	value = 0
}

func setZeroWithPointer(value *int) {
	*value = 0
}

func main() {
	//* Without pointer
	var value1 int = 10
	fmt.Println(value1)
	setZeroWithoutPointer(value1)
	fmt.Println(value1)

	//* With pointer
	var value2 int = 20
	fmt.Println(value2)
	setZeroWithPointer(&value2)
	fmt.Println(value2)

	//* convert (value to pointer & pointer to value)
	var value int = 100
	var pointer *int = &value
	var newValue int = (*pointer) + 1
	fmt.Println(newValue)

	/*
		10
		10

		20
		0

		101
	*/
}
