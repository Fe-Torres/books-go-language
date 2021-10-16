package main

import (
	"crud/database"
	"crud/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
