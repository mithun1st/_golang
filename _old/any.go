package main

import (
	"fmt"
	"reflect"
)

func main() {
	//1
	var v1 any = false
	fmt.Println(reflect.TypeOf(v1))
	v1 = "GoLang"
	fmt.Println(reflect.TypeOf(v1))

	//2
	var vSlice1 []any = []any{}
	vSlice1 = append(vSlice1, 101)
	vSlice1 = append(vSlice1, false)
	vSlice1 = append(vSlice1, 3.1416)
	fmt.Println(vSlice1)
	fmt.Println(reflect.TypeOf(vSlice1))

	//3
	var vSlice2 []any = make([]any, 0)
	vSlice2 = append(vSlice2, "newOne")
	vSlice2 = append(vSlice2, vSlice1...)
	fmt.Println(vSlice2)
	fmt.Println(reflect.TypeOf(vSlice2))
}
