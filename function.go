package main

import (
	"fmt"
)

func fnc() {
	fmt.Println("Hi There...")
}
func calculate(x int, y int) (int, int) {
	return (x + y), (x * y)
}

func main() {

	result1, _ := calculate(4, 6)
	r1, r2 := calculate(4, 6)
	fnc()

	fmt.Println(result1)
	fmt.Println(r1, " ", r2)
}
