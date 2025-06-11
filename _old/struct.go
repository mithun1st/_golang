package main

import (
	"fmt"
	"reflect"
)

// ---------------
type car struct {
	model  string
	color  string
	li_num int
}

// ---------------
type food struct {
	id        int
	food_name string
	cate      category
}
type category struct {
	cat_name  string
	cat_price float64
}

// ------------------
type laptop struct {
	model string
	price float64
	cpu   struct {
		cpu_type string
		core     int
	}
}

func main() {

	//regular struct
	var ferari = car{}
	ferari.color = "yellow"
	fmt.Println(ferari.color)

	bmw := car{
		model:  "AB",
		color:  "RED",
		li_num: 11,
	}

	var str string = "the car"
	str = str + ":" + fmt.Sprint(bmw.li_num)
	fmt.Println(str)
	fmt.Println(bmw)
	fmt.Println(reflect.TypeOf(bmw))

	//nested struct
	pizza := food{
		id:        101,
		food_name: "Cheese Pizza",
		cate: category{
			cat_name:  "italy",
			cat_price: 4.12,
		},
	}
	fmt.Println(pizza.id)
	fmt.Println(pizza.cate)
	fmt.Println(pizza.cate.cat_name)

	//anonimus struct
	student1 := struct {
		id   int
		name string
		dep  string
	}{
		dep: "cse",
	}
	fmt.Println(student1.id)
	fmt.Println(student1.name)
	fmt.Println(student1.dep)

	//nested
	laptop1 := laptop{
		model: "dell",
		price: 450.50,
		cpu: struct {
			cpu_type string
			core     int
		}{
			cpu_type: "intel",
			core:     5,
		},
	}
	fmt.Println(laptop1.model)
	fmt.Println(laptop1.cpu.cpu_type, " core: ", laptop1.cpu.core)

}
