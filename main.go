package main

import (
	"go-pgdb/handler"
)

func main() {
	handler.DefaultEnvironment()
	handler.HandleRequests()
}