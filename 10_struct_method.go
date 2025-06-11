package main

import "fmt"

type Student struct {
	roll int
	name string
}

// * Method
func (s Student) infoWithOutPointer(depName string) string {
	s.roll = -1
	return "Name: " + s.name + "\nDepartment: " + depName
}

func (s *Student) infoWithPointer(depName string) string {
	s.roll = -1
	return "Name: " + s.name + "\nDepartment: " + depName
}

func main() {
	//* Name struct
	var stu1 Student = Student{
		101,
		"stu1",
	}
	fmt.Println(stu1)
	fmt.Println(stu1.roll)
	fmt.Println(stu1.name)

	var stu2 Student
	fmt.Println(stu2)
	stu2.roll = 102
	stu2.name = "stu2"
	fmt.Println(stu2)

	//* Annonimus struct
	dep1 := struct {
		name   string
		credit float32

		techer struct {
			name   string
			rating float32
		}
	}{
		name:   "CSE",
		credit: 152.4,
		techer: struct {
			name   string
			rating float32
		}{
			name:   "Mr x",
			rating: 4.5,
		},
	}
	fmt.Println(dep1)

	//* Method call()
	stu3 := Student{201, "Mh Mithun"}
	fmt.Println(stu3)

	fmt.Println(stu3.infoWithOutPointer("EEE"))
	fmt.Println(stu3)
	fmt.Println(stu3.infoWithPointer("EEE"))
	fmt.Println(stu3)

	/*
		{101 stu1}
		101
		stu1
		{0 }
		{102 stu2}
		{CSE 152.4 {Mr x 4.5}}
		{201 Mh Mithun}
		Name: Mh Mithun
		Department: EEE
		{201 Mh Mithun}
		Name: Mh Mithun
		Department: EEE
		{-1 Mh Mithun}
	*/
}
