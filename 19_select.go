package main

import (
	"fmt"
	"sync"
	"time"
)

func process(c chan string, waitGroup *sync.WaitGroup, str string, delay int) {
	time.Sleep(time.Second * time.Duration(delay))
	c <- str

	waitGroup.Done()
}

func main() {

	var ch2 chan string = make(chan string, 1)
	var ch1 chan string = make(chan string, 1)

	var wg sync.WaitGroup

	//* Without Select
	wg.Add(2)
	go process(ch1, &wg, "Task 1", 2)
	go process(ch2, &wg, "Task 2", 1)
	wg.Wait()

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

	//* With Select
	wg.Add(2)
	go process(ch1, &wg, "Task X", 2)
	go process(ch2, &wg, "Task Y", 1)
	wg.Wait()

	for i := 0; i < 2; i++ {
		select {
		case p1 := <-ch1:
			fmt.Println(p1)
			break
		case p2 := <-ch2:
			fmt.Println(p2)
		}
	}

	/*
		Task 1
		Task 2
		Task Y
		Task X
	*/
}
