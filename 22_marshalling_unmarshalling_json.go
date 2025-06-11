package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string  `json:"name"`
	Roll int     `json:"roll"`
	Cgpa float32 `json:"cgpa"`
}

func main() {
	var stu1 student = student{Name: "Mahadi Hassan", Roll: 101, Cgpa: 3.14}
	fmt.Println(stu1)

	//* Marshal
	marshalJson, hasError1 := json.Marshal(stu1)
	fmt.Println(marshalJson, hasError1)
	fmt.Println(string(marshalJson))

	//* Unmarshal
	var stu2 student
	var hasError2 error = json.Unmarshal(marshalJson, &stu2)
	fmt.Println(stu2, hasError2)

	//* Unmarshal 2
	var stu3 student
	var hasError3 error = json.Unmarshal([]byte(`{"name":"Mithun","roll":102,"cgpa":3.14156}`), &stu3)
	fmt.Println(stu3, hasError3)

	/*
		{Mahadi Hassan 101 3.14}
		[123 34 110 97 109 101 34 58 34 77 97 104 97 100 105 32 72 97 115 115 97 110 34 44 34 114 111 108 108 34 58 49 48 49 44 34 99 103 112 97 34 58 51 46 49 52 125] <nil>
		{"name":"Mahadi Hassan","roll":101,"cgpa":3.14}
		{Mahadi Hassan 101 3.14} <nil>
		{Mithun 102 3.14156} <nil>
	*/
}
