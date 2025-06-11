package main

import (
	"fmt"
)

type employeFnc interface {
	getName() string
	getSalary(bonus int) float64
}

type employee struct {
	name       string
	salary     float64
	workingDay int
}

func (e employee) getName() string {
	return e.name
}

func (e employee) getSalary(b int) float64 {
	return (e.salary * float64(e.workingDay)) + float64(b)
}

func main() {

	per1 := employee{
		name:       "Mr x",
		salary:     2.01,
		workingDay: 12,
	}

	fmt.Println(per1.name)

	fmt.Println(employeFnc.getName(per1))
	fmt.Println((employeFnc.getSalary(per1, 5)))

}
