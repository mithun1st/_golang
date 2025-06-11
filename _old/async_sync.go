package main

import (
	"fmt"
	"time"
)

func asyncFnc() {
	// run parallel
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Hi... async")
	}()
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Bye... async")
	}()
	fmt.Println("async")
}

func syncFnc() {
	// run linear
	func() {
		time.Sleep(time.Second * 3)
		fmt.Println("Hi... sync")
	}()
	func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Bye... sync")
	}()
	fmt.Println("sync")
}

func bFnc() {
	time.Sleep(time.Second * 2)
	fmt.Println("There.....")
}

func main() {

	fmt.Println("start (for 10 sec)")

	asyncFnc()
	syncFnc()

	time.Sleep(time.Second * 10)
	fmt.Println("end")
}
