package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Services by Friends!")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("About: A community of service providers.")
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendString("Signup page coming soon...")
	})

	// Start server on port 8080
	log.Println("Frontend running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}