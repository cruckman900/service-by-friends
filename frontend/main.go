package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load HTML templates from ./templates directory
	engine := html.New("./templates", ".html")

	// Create a new Fiber app with template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static directories
	app.Static("/templates", "./templates")
	app.Static("/static", "./static")
	app.Static("/data", "./data")

    // Default route → index.html
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendFile("./templates/index.html")
    })

	// Start server on port 8080
	log.Println("Frontend running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}