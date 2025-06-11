package main

import (
	"fmt"
	myPackage "package_import_export/utils"
)

//* $go mod init package_import_export

func main() {
	fmt.Println(myPackage.NickName)
	myPackage.DisplayWithNickName(myPackage.NickName)

	/*
		Mithun
		Mahadi Hassan Mithun
	*/
}
