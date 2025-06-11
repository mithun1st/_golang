package main

import (
	"errors"
	"fmt"
)

func addditionAndDivision(a int, b int) (int, float32, error) {

	if b == 0 {
		return 0, 0, errors.New("Can't divided by Zoro")
	}
	div := float32(a) / float32(b)

	return (a + b), div, nil
}

func variadicFunction(nums ...int) int {
	var sum int = 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func closur() func(inc int) int {
	var i int = 0
	return func(incc int) int {
		i += incc
		return i
	}

}

func main() {
	//* Function
	add1, div1, err1 := addditionAndDivision(9, 5)
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {
		fmt.Println(add1, div1)
	}

	//* Errors
	add2, div2, err2 := addditionAndDivision(9, 0)
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println(add2, div2)
	}

	fmt.Println(variadicFunction(5, 5, 10))

	//* Closur
	var obj func(int) int = closur()
	fmt.Println(obj(100))
	fmt.Println(obj(2))
	fmt.Println(obj(1))

	/*
		14 1.8
		Can't divided by Zoro
		20
		100
		102
		103
	*/

}
