package main

import (
	"Google_sign_option/database"
	"Google_sign_option/route"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	app := fiber.New()
	route.SetupRoute(app)
	log.Fatal(app.Listen(":8080"))
}
