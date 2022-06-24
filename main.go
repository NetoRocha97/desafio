package main

import (
	"modulo/database"
	"modulo/server"
)

func main() {
	database.StartDb()
	server := server.NewServer()
	server.Run()
}
