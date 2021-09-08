package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pcpratheesh/golang-firebase-authentication-api/firebase"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	// initiate firebase
	firebaseAuth, err := firebase.InitiateClient()
	if err != nil {
		log.Fatal(err)
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("firebaseAuth", firebaseAuth)
		return c.Next()
	})

	app.Use(firebase.FirebaseAuthMiddleware)

	api := app.Group("api")
	v1 := api.Group("v1")

	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("connected")
	})
	app.Listen(":8081")

}
