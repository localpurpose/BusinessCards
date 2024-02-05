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

func SetupDefaultRoutes(app *fiber.App) {

	//------------------------------------------------
	// API
	//------------------------------------------------

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.Redirect("/1") // Redirect to the creator page
	})

	app.Get("/api/user/:userid", controller.GetDetails)

	app.Post("/api/user/create", controller.CreateUser)
	app.Post("/api/user/edit/:userid", controller.EditUser)
	app.Post("/api/user/delete/:userid", controller.DeleteUser)

	//Upload Image Handler.
	app.Post("/api/uploadimage", controller.UploadImage)

	//------------------------------------------------
	// Registration
	//------------------------------------------------
	app.Get("/createcard", controller.RenderRegister)
	app.Post("/createcard", controller.DoRegister)

	//------------------------------------------------
	// Main Application
	//------------------------------------------------
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/1") // Redirect to the creator page
	})

	app.Get("/:userid", controller.RenderUserProfile)

}
