package routes

import (
	"fmt"
	"github.com/buscard/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type (
	Host struct {
		Fiber *fiber.App
	}
)

func Setup(app *fiber.App) {

	//------------------------------------------------
	// app
	//------------------------------------------------

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("app")
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("api working")
	})

	app.Post("/api/user/create", controller.CreateUser)
	app.Post("/api/user/edit/:userid", controller.EditUser)
	app.Post("/api/user/delete/:userid", controller.DeleteUser)

	app.Get("/api/user/:userid", controller.GetDetails)

	app.Get("/:userid", func(c *fiber.Ctx) error {

		err := c.Render("pinkView/index", fiber.Map{
			"Title": "Brown",
		})
		if err != nil {
			log.Info(err)
		}
		return nil
	})
}
