package main

import (
	"gin-basic/cmd/server"
	"gin-basic/pkg/enum"
)

func main() {
	server.Run(enum.Development)
}
