package main

import (
	"flag"
	"fmt"
)

func main() {

	var name string
	var age int
	var isMale bool

	flag.StringVar(&name, "nameK", "Admin", "Enter name?")
	flag.IntVar(&age, "ageK", 1, "Enter age?")
	flag.BoolVar(&isMale, "sexK", false, "is male?")

	data := flag.String("a", "b", "c")

	flag.Parse()

	fmt.Println("Name:", name, "Age:", age, "isMale:", isMale)
	fmt.Println(*data)

	//$ go run 24_cli_flag.go -nameK=Mithun -ageK=1994 -sexK=true
	/*
		Name: Mithun Age: 1994 isMale: true
		b
	*/
}
