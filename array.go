package main

import (
	"fmt"
)

func main() {
	//------------------------
	var num1 [3]int = [3]int{1, 2}
	fmt.Println(num1)
	num1 = [3]int{11, 22, 33}
	fmt.Println(num1)

	//------------------------
	num2 := [4]bool{}
	num2[2] = true
	fmt.Println(num2)
	fmt.Println("len of array- ", len(num2))
	fmt.Println("cap of array- ", cap(num2))

}
