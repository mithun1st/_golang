package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

func main() {
	fmt.Println("Dot . . .")
	db, err := sql.Open("godror", "salesorder/ipiorder@210.4.75.155:1522/ORCL12")
	if err != nil {
		fmt.Println("#setup ", err.Error(), "#")
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("#not connected ", err.Error(), "#")

	} else {

		fmt.Println("#connected#")
	}
}
