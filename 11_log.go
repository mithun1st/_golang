package main

import (
	"fmt"
	"log"
	"log/slog"
)

func main() {

	appConfig := struct {
		id   int
		name string
	}{
		3, "Mh",
	}

	fmt.Println(appConfig)

	fmt.Printf("%v\n", appConfig)
	fmt.Printf("%+v\n", appConfig)
	fmt.Printf("%#v\n", appConfig)

	e := fmt.Errorf("basic error")
	fmt.Println(e)

	println("Hello world")

	slog.Info("a1")
	slog.Error("a2")

	log.Println("b1")

	log.Fatalln("exit 1")
	log.Panicln("halt")

	/*
		{3 Mh}
		{3 Mh}
		{id:3 name:Mh}
		struct { id int; name string }{id:3, name:"Mh"}

		basic error

		Hello world

		2025/07/31 23:44:21 INFO a1
		2025/07/31 23:44:21 ERROR a2

		2025/07/31 23:44:21 b1

		2025/07/31 23:44:21 exit 1
		exit status 1
	*/

}
