package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//* String
	var str string = "Mithūn"
	fmt.Println(str)
	fmt.Println("Length of", str, "is", utf8.RuneCountInString(str))

	for i, v := range str {
		fmt.Println(i, v)
	}
	for i, v := range str {
		fmt.Println(i, v, string(v))
	}

	//* Rune
	var str2 []rune = []rune("Mithūn")
	fmt.Println(str2)
	fmt.Println("Length of", str2, "is", len(str2))

	for i2, v2 := range str2 {
		fmt.Println(i2, v2, string(v2))
	}

	//* Multiline
	var str3 string = `Mahadi Hassan
Mithun`
	fmt.Println(str3)

	/*
		Mithūn
		Length of Mithūn is 6
		0 77
		1 105
		2 116
		3 104
		4 363
		6 110
		0 77 M
		1 105 i
		2 116 t
		3 104 h
		4 363 ū
		6 110 n
		[77 105 116 104 363 110]
		Length of [77 105 116 104 363 110] is 6
		0 77 M
		1 105 i
		2 116 t
		3 104 h
		4 363 ū
		5 110 n
		Mahadi Hassan
		Mithun
	*/
}
