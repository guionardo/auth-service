package main

import (
	"fmt"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email youremail@provider.com
// @host localhost:3000
// @BasePath /
func main() {
	fmt.Println("Auth Service - golang version")
	// Create new Fiber application
	app := createServer()

	// Listen on the port '3000'
	app.Listen(":3000")
}
