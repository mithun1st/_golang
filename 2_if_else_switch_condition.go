package main

import (
	"fmt"
)

func main() {

	//* If Else
	var mark int = 30

	if mark < 33 {
		fmt.Println("Fail")
	} else if mark <= 100 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Wrong")
	}

	//* Switch
	var chh rune = 'a'

	switch chh {
	case 'a':
		fmt.Println("Apple")
		break
	case 'b':
		fmt.Println("Ball")
		break
	case 'c':
		fmt.Println("Cat")
		break
	}

	//* Table
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	/*
		Fail
		Apple
		true
		true
		true
		false
		true
		false
		false
		false
	*/
}
