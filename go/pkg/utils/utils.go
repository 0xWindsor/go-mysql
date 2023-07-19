package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New()

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
