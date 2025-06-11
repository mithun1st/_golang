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

	//switch
	var value int = 4
	switch value {
	case 1:
		fmt.Println("a")
		break
	case 4:
		fmt.Println("b")
		break
	default:
		fmt.Println("c")
		break
	}
}
