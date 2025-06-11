package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchWeater(city string) string {
	time.Sleep(time.Second)
	var result string = fmt.Sprint(city, ": ", int(city[0]))
	return result
}

func fetchWeaterWithChannel(city string, ch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	var result string = fmt.Sprint(city, ": ", int(city[0]))
	ch <- result
	wg.Done()
}

func main() {
	var cityes []string = []string{"Dhaka", "Natore", "Rajshahi"}

	//* Regular
	timeStart1 := time.Now()
	for _, data := range cityes {
		var re = fetchWeater(data)
		fmt.Println(re)
	}
	fmt.Println("Time Cost1:", time.Since(timeStart1))

	//* Channel
	timeStart2 := time.Now()

	var ch chan string = make(chan string, len(cityes))
	var wg sync.WaitGroup

	for _, data := range cityes {
		wg.Add(1)
		go fetchWeaterWithChannel(data, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for data := range ch {
		fmt.Println(data)
	}

	fmt.Println("Time Cost2:", time.Since(timeStart2))

	/*
		Dhaka: 68
		Natore: 78
		Rajshahi: 82
		Time Cost1: 3.008606084s

		Rajshahi: 82
		Dhaka: 68
		Natore: 78
		Time Cost2: 1.002693792s

	*/
}
