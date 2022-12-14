package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golangProxy/controllers/pingController"
	"golangProxy/controllers/transparencyPortalController"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	app := fiber.New()

	app.Get("/ping", pingController.Index)

	app.Get("/transparency_portal/:id", transparencyPortalController.Show)

	err := app.Listen(":8001")

	if err != nil {
		log.Fatal(err)
	}
}
