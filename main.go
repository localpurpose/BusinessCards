package main

import (
	db "github.com/buscard/config"
	"github.com/buscard/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"

	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	db.Connect()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	//Static Views Files
	app.Static("user/assets/brownView/", "./views/brownView/assets")
	app.Static("user/assets/blueView/", "./views/blueView/assets")
	app.Static("user/assets/orangeView/", "./views/orangeView/assets")
	app.Static("user/assets/pinkView/", "./views/pinkView/assets")
	app.Static("user/assets/qrs/", "./usersData")

	//Static Main Page files
	app.Static("assets/main/", "./views/home/assets")

	//Static Registration files
	app.Static("assets/reg/", "./views/registration/assets")

	routes.SetupDefaultRoutes(app)

	//------------------------------------------------
	// 404 Middleware
	//------------------------------------------------
	app.Use(func(c *fiber.Ctx) error {
		return c.Render("home/404", nil) // => 404 "Not Found"
	})

	m := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("visitkabot.ru"),
		Cache:      autocert.DirCache("./certs"),
	}

	cfg := &tls.Config{
		// Get Certificate from Let's Encrypt
		GetCertificate: m.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}

	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(app.Listener(ln))
}
