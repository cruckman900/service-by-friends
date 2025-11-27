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

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Services by Friends",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{
			"Title": "About Us",
		})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("signup", fiber.Map{
			"Title": "Sign Up",
		})
	})

	// Start server on port 8080
	log.Println("Frontend running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}