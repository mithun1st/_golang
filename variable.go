package main

import (
	"fmt"
)

func main() {

	//-------------------basic variable

	var isMale bool = true
	var num int = -94 //default 64
	var str string = "MH Mithun"
	var pi float64 = 3.1416
	var guest uint = 3 //default 64
	var rgb byte = 255

	fmt.Println(isMale)
	fmt.Println(num)
	fmt.Println(str)
	fmt.Println(pi)
	fmt.Println(guest)
	fmt.Println(rgb)

	fmt.Print("is male ", isMale, ", the num is ", num, "\n")
	fmt.Printf("bool: %t, int: %d, string: %v, float: %.2f, unit: %d, byte: %d \n",
		isMale, num, str, pi, guest, rgb)
	fmt.Printf("bool: %T, int: %T, string: %T, float: %T, unit: %T, byte: %T \n",
		isMale, num, str, pi, guest, rgb)

	//dynamic assign
	var1 := 255
	var2 := 6.5
	var3 := "hi there"
	var4 := false
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	// fmt.Println(reflect.TypeOf(var1))
	// fmt.Println(reflect.TypeOf(var2))
	// fmt.Println(reflect.TypeOf(var3))
	// fmt.Println(reflect.TypeOf(var4))

	//-------------------for more specify size

	var num1 int8 = 127
	var num2 int16 = 32767
	var num3 int32 = 2147483647

	var pi1 float32 = 1.0

	var guest1 uint8 = 255
	var guest2 uint32 = 1000000000

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(pi1)
	fmt.Println(guest1)
	fmt.Println(guest2)
}
