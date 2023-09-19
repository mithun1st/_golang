package main

import (
	"fmt"
)

func main() {

	var num int = 9
	winner := true

	if num > 0 {
		fmt.Println("first block")
	} else if winner {
		fmt.Println("second block")
	} else {
		fmt.Println("all condition fail")
	}

	if num > 0 && winner {
		fmt.Println("AND Operation")
	}

	if (num < 0) || winner {
		fmt.Println("OR Operation")
	}

}
