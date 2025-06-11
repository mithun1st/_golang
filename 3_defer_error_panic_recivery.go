package main

import (
	"errors"
	"fmt"
)

func errorExample() {
	var err3 error = fmt.Errorf("Error 1")
	var err4 error = errors.New("Error 2")
	fmt.Println(err3, err3.Error())
	fmt.Println(err4, err4.Error())
}

func deferExample() {
	// defer will run after function end
	fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	fmt.Println("Defer 3")
}

func recoveryExample() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err, "(Panic error handel)")
		}
	}()
	fmt.Println("Recovery 1")
	panic("Recovery 2")
	fmt.Println("Recovery 3")
}

func panicExample() {
	// nothing will execute after panic if not recovery
	panic("Panic Error")
	fmt.Println("This line will not print if panic error") //dead code
}

func main() {

	//* 1 Error
	errorExample()

	//* 2 Defer
	deferExample()

	//* 3 Recovery
	recoveryExample()
	fmt.Println("Recovery End")

	//* 4 Panic
	panicExample()
	fmt.Println("Panic End")

	/*
		Error 1 Error 1
		Error 2 Error 2

		Defer 1
		Defer 3
		Defer 2

		Recovery 1
		Recovery 2 (Panic error handel)
		Recovery End

		panic: Panic Error

		goroutine 1 [running]:
		main.panicExample(...)
			/Volumes/dev-space/source-code/go/3_defer_error_panic_recivery.go:36
		main.main()
			/Volumes/dev-space/source-code/go/3_defer_error_panic_recivery.go:53 +0x69
		exit status 2
	*/
}
