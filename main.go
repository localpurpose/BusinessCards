package main

import (
	db "github.com/buscard/config"
	"github.com/buscard/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {

	db.Connect()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	app.Static("assets/brownView/", "./views/brownView/assets")
	app.Static("assets/blueView/", "./views/blueView/assets")
	app.Static("assets/orangeView/", "./views/orangeView/assets")
	app.Static("assets/pinkView/", "./views/pinkView/assets")
	app.Static("assets/qrs/", "./usersData")

	routes.Setup(app)
	log.Fatal(app.Listen(":80"))
}
