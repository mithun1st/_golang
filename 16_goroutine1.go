package main

import (
	"fmt"
	"sync"
	"time"
)

func task(taskName string) {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(taskName, i)
	}
}

func taskWithWaitGroup(taskName string, lenght int, wg *sync.WaitGroup) {
	for i := 1; i <= lenght; i++ {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(taskName, i)
	}
	wg.Done()
}

func main() {
	//* Without Goroutine
	task("Task A:")
	task("Tast B:")

	//* With Goroutine
	go task("Task M (goroutine):")
	go task("Task N (goroutine):")
	time.Sleep(time.Second * 2)

	//* Wait Group
	var wg sync.WaitGroup

	wg.Add(1)
	go func(w *sync.WaitGroup) {
		time.Sleep(time.Second)
		fmt.Println("Task")
		w.Done()
	}(&wg)
	wg.Wait()

	wg.Add(2)
	go taskWithWaitGroup("Task X", 3, &wg)
	go taskWithWaitGroup("Task Y", 5, &wg)
	wg.Wait()

	/*
		Task A: 1
		Task A: 2
		Task A: 3
		Tast B: 1
		Tast B: 2
		Tast B: 3
		Task M (goroutine): 1
		Task N (goroutine): 1
		Task M (goroutine): 2
		Task N (goroutine): 2
		Task N (goroutine): 3
		Task M (goroutine): 3
		Task
		Task Y 1
		Task X 1
		Task X 2
		Task Y 2
		Task Y 3
		Task X 3
		Task Y 4
		Task Y 5
	*/
}
