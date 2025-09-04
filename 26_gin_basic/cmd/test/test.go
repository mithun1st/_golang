package main

import (
	"fmt"
	logger "gin-basic/pkg/logging"
	jwttoken "gin-basic/pkg/token"
	"time"
)

func main() {

	now := time.Date(2025, 10, 21, 0, 0, 0, 0, time.UTC)
	exp := time.Date(2025, 10, 22, 0, 0, 0, 0, time.UTC)

	re := exp.After(now)
	logger.Info(re)

	unix := (time.Now().Unix())
	unixU := (time.Now().UTC().Unix())

	fmt.Println(unix)
	fmt.Println(unixU)

	fmt.Println(time.Unix(unix, 0))
	fmt.Println(time.Unix(unixU, 0))

	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())

	// ss := []byte("your-secret-key")
	// fmt.Println(string(ss))
	logger.Info("hello")
	logger.Error("world")

	var key string = ""

	mp1 := map[string]any{
		"user": "mithun.2121@yahoo.com",
		"exp":  time.Now().Add(time.Hour * 24 * 90).Unix(),
	}
	token1, _ := jwttoken.Encript(mp1, key)
	logger.Info(token1)

	var t string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoibWl0aHVuLjIxMjFAeWFob28uY29tIn0.UnB1CVFo1QcFU6RkK5sxlRr_bH1RVnEZoktXDZABSzo"
	isValid := jwttoken.IsTokenValid(t, "")
	logger.Info(isValid)

	data, err := jwttoken.Decript(t, "")
	logger.Info(err)
	logger.Info(data)
	ss := "Bearar:token file"
	fmt.Println(ss[7:])
}
