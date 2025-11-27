package main

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Provider document model
type Provider struct {
	ID        string    `json:"id"`        // use doc ID as string
	Name      string    `json:"name"`
	Service   string    `json:"service"`
	CreatedAt time.Time `json:"createdAt"`
}

func initFirebaseApp() *firebase.App {
	ctx := context.Background()

	// Read JSON from env var
	credsB64 := os.Getenv("FIREBASE_KEY_BASE64")
	credsJSON, err := base64.StdEncoding.DecodeString(credsB64)
	if err != nil {
		log.Fatalf("failed to decode FIREBASE_KEY_BASE64: %v", err)
	}
	opt := option.WithCredentialsJSON(credsJSON)

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "services-by-friends",
	}, opt)
	if err != nil {
		log.Fatalf("firebase app init failed: %v", err)
	}
	return app
}

func initFirestore(app *firebase.App) *firestore.Client {
	ctx := context.Background()
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("firestore client init failed: %v", err)
	}
	return client
}

func main() {
	// Initialize Firebase Admin + Firestore
	app := initFirebaseApp()
	db := initFirestore(app)
	defer db.Close()

	// Fiber app
	api := fiber.New()

	// GET /providers → list providers
	api.Get("/providers", func(c *fiber.Ctx) error {
		ctx := context.Background()
		iter := db.Collection("providers").OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
		defer iter.Stop()

		var list []Provider
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
			var p Provider
			if err := doc.DataTo(&p); err != nil {
				return c.Status(500).SendString(err.Error())
			}
			// Ensure the doc ID is set
			p.ID = doc.Ref.ID
			list = append(list, p)
		}
		return c.JSON(list)
	})

	// POST /providers → add provider
	// Body: { "name": "Robert Long", "service": "Tattooing" }
	api.Post("/providers", func(c *fiber.Ctx) error {
		var input struct {
			Name    string `json:"name"`
			Service string `json:"service"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).SendString("invalid body: " + err.Error())
		}
		if input.Name == "" || input.Service == "" {
			return c.Status(422).SendString("name and service are required")
		}

		ctx := context.Background()
		docRef, _, err := db.Collection("providers").Add(ctx, map[string]interface{}{
			"Name":      input.Name,
			"Service":   input.Service,
			"CreatedAt": time.Now(),
		})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(fiber.Map{"id": docRef.ID})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // fallback for local dev
	}
	log.Println("Providers API with Firestore running on http://localhost:" + port)
	log.Fatal(api.Listen(":" + port))
}