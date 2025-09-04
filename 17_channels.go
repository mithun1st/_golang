package main

import (
	"fmt"
	"time"
)

func main() {

	//* Channel
	var ch1 chan string = make(chan string, 0)

	go func() {
		ch1 <- "ping A1"
		ch1 <- "ping A2"
		ch1 <- "ping A3"
	}()

	var resultA1 string = <-ch1
	var resultA2 string = <-ch1
	var resultA3 string = <-ch1
	close(ch1)

	fmt.Println(resultA1)
	fmt.Println(resultA2)
	fmt.Println(resultA3)
	fmt.Println(ch1)

	//* Buffer
	var ch2 chan string = make(chan string, 2)

	ch2 <- "ping B1"
	ch2 <- "ping B2"

	var resultB1 string = <-ch2
	var resultB2 string = <-ch2

	fmt.Println(resultB1)
	fmt.Println(resultB2)

	// Data can be insert again after channel being empty

	ch2 <- "ping reinsert 1"
	ch2 <- "ping reinsert 2"

	/*
		ch2 <- "ping reinsert 3" => fatal error: all goroutines are asleep - deadlock!
	*/

	go func() {
		// insert in another thread to prevent deadlock
		ch2 <- "ping reinsert 3"
	}()

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// for data := range ch2 {
	// 	fmt.Println(data)
	// }
	close(ch2)

	//* Synchronization Passing
	var chSender chan string = make(chan string, 1)
	var chReceiver chan string = make(chan string, 1)

	chSender <- "ping passing"

	go func(s chan string, r chan string) {
		var data string = <-s
		r <- data
	}(chSender, chReceiver)

	fmt.Println(<-chReceiver)
	time.Sleep(time.Second)

	/*
		ping A1
		ping A2
		ping A3
		0xc000076070
		ping B1
		ping B2
		ping reinsert 1
		ping reinsert 2
		ping reinsert 3
		ping passing
	*/
}
