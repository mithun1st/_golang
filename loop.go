package main

import (
	"fmt"
)

func main() {

	//for loop
	for i := 1; i < 5; i++ {
		fmt.Println("for loop ", i)
	}

	//while loop
	var count int = 10
	for count > 0 {
		fmt.Println("while loop ", count)
		count--
	}

	count = 10
	for {
		fmt.Println("bla bla ", count)
		count++
		if count > 20 {
			break
		}
	}

	for count = 1; count <= 15; count++ {
		if count%2 == 0 {
			continue
		}
		fmt.Println(count)
	}

}
