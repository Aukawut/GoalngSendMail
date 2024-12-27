package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mail/mail"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", mail.SendMail)

	log.Fatal(app.Listen(":3005"))
}
