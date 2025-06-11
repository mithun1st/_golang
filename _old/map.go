package main

import (
	"fmt"
)

func main() {
	var mp map[string]int
	mp = map[string]int{
		"key1": 101,
		"key2": 202,
		"key3": 303,
	}
	mp["key9"] = 404
	fmt.Println(len(mp))
	fmt.Println(mp)
	fmt.Println(mp["key9"])

	//check key
	v, has := mp["key9"]
	fmt.Println(v, has)
	v, has = mp["key10"]
	fmt.Println(v, has)

	for k, v := range mp {
		fmt.Println(k, " ", v)
	}

	//delete
	delete(mp, "key9")
	fmt.Println(mp)

}
