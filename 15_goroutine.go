package main

import (
	"fmt"
	"time"
)

func delayDisplay(num int) {
	fmt.Println("Display:", num)
	time.Sleep(time.Second)
}

func main() {

	delayDisplay(1)
	delayDisplay(2)
	go delayDisplay(3) // run in thread
	delayDisplay(4)

	// system delay (otherwise main() end before goroutine end)
	time.Sleep(time.Second * 3)
	fmt.Println("Main END")

	/*
		Display: 1
		Display: 2
		Display: 4
		Display: 3
		END
	*/
}
