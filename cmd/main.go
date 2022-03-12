package main

import (
	"crud/internal/database"
	"crud/internal/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
