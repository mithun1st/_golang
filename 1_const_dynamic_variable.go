package main

import (
	"fmt"
	"reflect"
)

func main() {

	//* Const
	const pi float32 = 3.14156

	//* Print
	fmt.Print(pi)
	fmt.Printf("Value:%.2f Type:%T\n", pi, pi)
	fmt.Println("PI value:", pi)

	//* Initial Type
	const (
		piWith2decPoint float32 = 3.14
		mh              string  = "Mahadi Hassan"
	)
	var (
		num1 int = 1
		num2 int = 2
	)

	var num3 int = 3

	num4 := 4

	num5, num6 := 5, 6

	var num7 int
	num7 = 7

	fmt.Println(num1, num2, num3, num4, num5, num6, num7)

	//* Datatype
	var var0 int64 = 123456789 //long int
	var var1 int = -123456789  //same as int64 bit
	var var2 int8 = -128       //-128 to 127
	var var3 uint8 = 255       //0 to 255 (only take positive number)

	var var4 int16 = -32768
	var var5 int32 = -2147483648

	var var6 float32 = 3.14
	var var7 float64 = 3.141569
	var var8 byte = 255 //same as int8

	var var9 bool = true
	var var10 rune = 'M'

	fmt.Println(var0)
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)
	fmt.Println(var7)
	fmt.Println(var8)
	fmt.Println(var9)
	fmt.Println(string(var10))

	//* Dynamic variable
	var dynamic1 any = 12
	fmt.Println("Value:", dynamic1, "Type:", reflect.TypeOf(dynamic1))

	dynamic1 = true
	fmt.Println("Value:", dynamic1, "Type:", reflect.TypeOf(dynamic1))

	dynamic1 = "Mahadi Hassan"
	fmt.Println("Value:", dynamic1, "Type:", reflect.TypeOf(dynamic1))

	/*
		3.14156Value:3.14 Type:float32
		PI value: 3.14156
		1 2 3 4 5 6 7
		123456789
		-123456789
		-128
		255
		-32768
		-2147483648
		3.14
		3.141569
		255
		true
		M
		Value: 12 Type: int
		Value: true Type: bool
		Value: Mahadi Hassan Type: string
	*/
}
