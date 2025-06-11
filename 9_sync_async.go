package main

import (
	"fmt"
	"time"
)

// * Sync
func sync(callBack func()) {
	fmt.Println("sync start")

	fmt.Println("sync wait for 1...")
	func() {
		time.Sleep(time.Second * 2)
		fmt.Println("(sync 1)")
		time.Sleep(time.Second * 1)
	}()

	fmt.Println("sync wait for 2...")
	func() {
		time.Sleep(time.Second * 2)
		fmt.Println("(sync 2)")
		time.Sleep(time.Second * 1)
	}()

	fmt.Println("sync end")
	callBack()
}

// * Async
func async() {
	fmt.Println("async start")

	fmt.Println("async wait for 1...")
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("(async 1)")
		time.Sleep(time.Second * 1)
	}()

	fmt.Println("async wait for 2...")
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("(async 2)")
		time.Sleep(time.Second * 1)
	}()

	fmt.Println("async end")
}

func main() {
	sync(func() {
		fmt.Println("Call back func call()")
	})
	async()

	// This line hold the main() untill the async() finish
	time.Sleep(time.Second * 7)
	fmt.Println("Main Program End")

	/*
	   sync start
	   sync wait for 1...
	   (sync 1)
	   sync wait for 2...
	   (sync 2)
	   sync end
	   Call back func call()
	   async start
	   async wait for 1...
	   async wait for 2...
	   async end
	   (async 2)
	   (async 1)
	   Main Program End
	*/
}
