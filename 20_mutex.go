package main

import (
	"fmt"
	"sync"
)

func incrementWithoutMutex(c *int, w *sync.WaitGroup) {
	*c++
	w.Done()
}

func incrementWithMutex(c *int, w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*c++
	m.Unlock()

	w.Done()
}

func main() {

	var wg sync.WaitGroup

	//* Without Mutex
	var counter1 int = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementWithoutMutex(&counter1, &wg)
	}
	wg.Wait()
	fmt.Println(counter1)

	//* With Mutex
	var counter2 int = 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementWithMutex(&counter2, &wg, &mu)
	}
	wg.Wait()
	fmt.Println(counter2)

	/*
		873
		1000
	*/

}
