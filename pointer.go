package main

import (
	"fmt"
	"reflect"
)

func main() {
	var pi float64 = 3.1415
	var pointer *float64 = &pi

	fmt.Println(pi)
	fmt.Println(pointer)

	fmt.Println(reflect.TypeOf(pointer))

	fmt.Println(*pointer)
}
