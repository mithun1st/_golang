package main

import (
	"fmt"
)

func main() {

	//* For
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		if i == 8 {
			fmt.Printf("Loop Break At %d count\n", i)
			break
		}
		fmt.Println("For:", i)
	}

	//* While
	var i int = 0
	for i > -10 {
		fmt.Println("While", i)
		i--
	}

	//* Infinity
	var j int = 0
	for {
		fmt.Print(j, ", ")
		if j == 7 {
			break
		}
		j++
	}
	fmt.Println()

	/*
		For: 0
		For: 1
		For: 2
		For: 3
		For: 4
		For: 6
		For: 7
		Loop Break At 8 count
		While 0
		While -1
		While -2
		While -3
		While -4
		While -5
		While -6
		While -7
		While -8
		While -9
		0, 1, 2, 3, 4, 5, 6, 7,
	*/
}
