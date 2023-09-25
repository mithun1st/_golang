package main

import (
	"errors"
	"fmt"
)

func getUser(id int) (string, error) {
	if id < 10 {
		return "found", nil
	} else {
		return "", errors.New("not_found")
	}
}

func main() {
	emp1, isError := getUser(8) //lets assume- id of below 10 is valid

	if isError != nil {
		fmt.Println(isError.Error())
	} else {
		fmt.Println(emp1)
	}

}
