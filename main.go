package main

import (
	db "github.com/buscard/config"
	"github.com/buscard/routes"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	db.Connect()

	app := fiber.New()

	routes.Setup(app)
	log.Fatal(app.Listen(":80"))
}
