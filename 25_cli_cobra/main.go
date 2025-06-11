package main

import (
	cmd "cli-cobra/command"
)

func main() {
	cmd.Execute()

	//? go run main.go == cal (after build CLI "main.go" replace with "cal")
	/*
		$ go run main.go -h
		A Fast and Flexible Static Site Generator LONG

		Usage:
		  cal [command]

		Available Commands:
		  completion  Generate the autocompletion script for the specified shell
		  help        Help about any command
		  mh          Mh SHORT
		  pi          PI SHORT

		Flags:
		  -h, --help   help for cal

		Use "cal [command] --help" for more information about a command.
	*/

	//* mh [command]
	/*
		$ go run main.go mh -h
		Mh LONG

		Usage:
		  cal mh [flags]

		Flags:
		      --fullname   First name, Last name, Nick nmae
		  -h, --help       help for mh

	*/

	/*
		$ go run main.go mh
		MH Mithun

		$ go run main.go mh --fullname
		Mahadi Hassan Mithun
	*/

	//* pi [command]
	/*
		$ go run main.go pi -h
		PI LONG

		Usage:
		  cal pi [flags]

		Flags:
		  -h, --help           help for pi
		  -p, --plus int       Plus number with pi
		  -s, --short          View in short form
		      --title string   Title of pi (default "PI Value")
	*/

	/*
		$ go run main.go pi
		PI Value : 3.141592653589793

		$ go run main.go pi -p 2
		5.141592653589793
		$ go run main.go pi --plus 3
		6.141592653589793

		$ go run main.go pi -s
		3.14
		$ go run main.go pi --short
		3.14

		$ go run main.go pi --title ValueOfPi
		ValueOfPi : 3.141592653589793
	*/

}
