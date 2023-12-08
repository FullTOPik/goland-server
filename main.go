package main

import (
	database "server/internal/config"
	"server/internal/controllers/userController"
	"server/pkg/server"
)

func main() {
	database.Connect()
	defer database.Disconnect()

	server.Use("/user", userController.GetRouters())

	server.Listen(5000, nil)
}
