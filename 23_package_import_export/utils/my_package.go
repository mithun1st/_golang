package myPackage

import (
	"fmt"
)

//*	Started with lowercase declare as a PRIVATE variable

const firstName string = "Mahadi"
const lastName string = "Hassan"
const NickName string = "Mithun"

func getFullName() string {
	return fmt.Sprint(firstName, " ", lastName)
}

func DisplayWithNickName(nickName string) {
	fmt.Println(getFullName(), nickName)
}
