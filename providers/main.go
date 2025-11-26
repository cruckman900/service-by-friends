package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Provider struct {
	ID          int 	`json:"id"`
	Name        string 	`json:"name"`
	Service		string 	`json:"service"`
}

var providers = []Provider{
	{ID: 1, Name: "Robert Long", Service: "Tattooing"},
	{ID: 2, Name: "Joe Sparkman", Service: "Custom Guitars"},
}

func main() {
	app := fiber.New()

	// GET /providers -> returns JSON list
	app.Get("/providers", func(c *fiber.Ctx) error {
		return c.JSON(providers)
	})

	// Start server on port 8090
	log.Println("Providers service running on http://localhost:8081")
	log.Fatal(app.Listen(":8081"))
}