package main

import (
	"fmt"
	"reflect"
)

func anyExample() {
	var var1 any = "Mh Mithun"
	fmt.Println(var1)

	var var2 any = 123
	fmt.Println(var2)

	var2 = 3.141569
	fmt.Println(var2)

	var2 = false
	fmt.Println(var2)

	var2 = []string{"Mahadi", "Hassan"}
	fmt.Println(var2)
}

func genericsExample[T1 any](value T1) T1 {
	fmt.Println(reflect.TypeOf(value))
	return value
}

func main() {
	//* Any
	anyExample()

	//* Generics
	fmt.Println(genericsExample[int8](1))
	fmt.Println(genericsExample[int16](2))
	fmt.Println(genericsExample[int32](3))
	fmt.Println(genericsExample(4))

	/*
		Mh Mithun
		123
		3.141569
		false
		[Mahadi Hassan]

		int8
		1
		int16
		2
		int32
		3
		int
		4
	*/
}
