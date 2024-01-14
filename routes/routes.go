package routes

import (
	"github.com/buscard/controller"
	"github.com/gofiber/fiber/v2"
)

type (
	Host struct {
		Fiber *fiber.App
	}
)

func Setup(app *fiber.App) {

	//------------------------------------------------
	// API
	//------------------------------------------------

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/1") // Redirect to the creator page
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Redirect("/1") // Redirect to the creator page
	})

	app.Post("/api/user/create", controller.CreateUser)
	app.Post("/api/user/edit/:userid", controller.EditUser)
	app.Post("/api/user/delete/:userid", controller.DeleteUser)

	app.Get("/api/user/:userid", controller.GetDetails)

	//------------------------------------------------
	// Main Application
	//------------------------------------------------

	app.Get("/:userid", controller.RenderUserProfile)
}
