package main

import (
	"books-api/app"
	"books-api/pkg/common"
)

func main() {
	app.StartServer().Run(common.ServerPort)
}
