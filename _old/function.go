package main

import (
	"errors"
	"fmt"
)

func fnc() {
	fmt.Println("Hi There...")
}

func calculate(x int, y int) (int, int) {
	var sum int = x + y
	var mul int = x * y
	return sum, mul
}

func errorCheck() (bool, error, error) {
	return false, nil, errors.New("this is an error!")
}

func main() {

	fnc()

	result1, _ := calculate(4, 6)
	r1, r2 := calculate(4, 6)

	eBool, e1, e2 := errorCheck()

	fmt.Println(result1)
	fmt.Println(r1, " ", r2)

	fmt.Println(eBool, "\t", e1, "\t", e2)
}
