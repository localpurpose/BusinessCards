package routes

import (
	"fmt"
	"github.com/buscard/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
)

type (
	Host struct {
		Fiber *fiber.App
	}
)

func Setup(app *fiber.App) {

	hosts := map[string]*Host{}

	api := fiber.New()
	hosts["api.localhost"] = &Host{api}

	//------------------------------------------------
	// Main domain
	//------------------------------------------------

	main := fiber.New()
	hosts["localhost"] = &Host{main}

	main.Get("/1", func(c *fiber.Ctx) error {
		engine := html.New("./views/brownView", ".html")
		main.Static("assets/", "./views/brownView/assets")
		main = fiber.New(fiber.Config{
			Views: engine,
		})
		return c.Render("idex", ".html")
	})

	//------------------------------------------------
	// API
	//------------------------------------------------

	api.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("API")
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	api.Post("/user/create", controller.CreateUser)
	api.Post("/user/edit/:userid", controller.EditUser)
	api.Post("/user/delete/:userid", controller.DeleteUser)

	api.Get("/user/:userid", controller.GetDetails)

	app.Use(func(c *fiber.Ctx) error {
		host := hosts[c.Hostname()]
		log.Info(hosts[c.Hostname()])
		if host == nil {
			return c.SendStatus(fiber.StatusNotFound)
		} else {
			host.Fiber.Handler()(c.Context())
			return nil
		}
	})

}
